package redisx

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

////////////////////////////////////////////////////////////
// -------------------- Interface ------------------------
////////////////////////////////////////////////////////////

type DataSource interface {
	WithTTL(ttl time.Duration) DataSourceBuilder

	// Record CRUD (stored as JSON string)
	Create(ctx context.Context, key string, value any) error
	Update(ctx context.Context, key string, value any) error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string, dest any) error

	// Index operations
	AddToIndex(ctx context.Context, indexKey, id string, score float64) error
	RemoveFromIndex(ctx context.Context, indexKey, id string) error
	GetAllFromIndex(ctx context.Context, indexKey string, start, stop int64) ([]string, error)

	// Atomic helpers (production-safe)
	UpsertWithIndex(ctx context.Context, key, indexKey, id string, score float64, value any, ttl time.Duration) error
	DeleteWithIndex(ctx context.Context, key, indexKey, id string) error

	// Bulk read
	GetAll(ctx context.Context, indexKey string, destSlice any) error
}

type DataSourceBuilder interface {
	Create(ctx context.Context, key string, value any) error
	Update(ctx context.Context, key string, value any) error
	AddToIndex(ctx context.Context, indexKey, id string, score float64) error
	UpsertWithIndex(ctx context.Context, key, indexKey, id string, score float64, value any) error
}

////////////////////////////////////////////////////////////
// -------------------- Struct ---------------------------
////////////////////////////////////////////////////////////

type dataSource struct {
	RDb *redis.Client
}

type dataSourceBuilder struct {
	parent *dataSource
	ttl    time.Duration
}

////////////////////////////////////////////////////////////
// -------------------- Compile-time checks --------------
////////////////////////////////////////////////////////////

var _ DataSource = (*dataSource)(nil)
var _ DataSourceBuilder = (*dataSourceBuilder)(nil)

////////////////////////////////////////////////////////////
// -------------------- Constructor ----------------------
////////////////////////////////////////////////////////////

func NewDataSource(rdb *redis.Client) DataSource {
	return &dataSource{
		RDb: rdb,
	}
}

////////////////////////////////////////////////////////////
// -------------------- Builder --------------------------
////////////////////////////////////////////////////////////

func (r *dataSource) WithTTL(ttl time.Duration) DataSourceBuilder {
	return &dataSourceBuilder{
		parent: r,
		ttl:    ttl,
	}
}

func (b *dataSourceBuilder) Create(ctx context.Context, key string, value any) error {
	return b.parent.saveJSON(ctx, key, value, b.ttl)
}

func (b *dataSourceBuilder) Update(ctx context.Context, key string, value any) error {
	return b.parent.saveJSON(ctx, key, value, b.ttl)
}

func (b *dataSourceBuilder) AddToIndex(
	ctx context.Context,
	indexKey, id string,
	score float64,
) error {
	return b.parent.RDb.ZAdd(ctx, indexKey, redis.Z{
		Score:  score,
		Member: id,
	}).Err()
}

func (b *dataSourceBuilder) UpsertWithIndex(
	ctx context.Context,
	key, indexKey, id string,
	score float64,
	value any,
) error {
	return b.parent.UpsertWithIndex(ctx, key, indexKey, id, score, value, b.ttl)
}

////////////////////////////////////////////////////////////
// -------------------- Core CRUD ------------------------
////////////////////////////////////////////////////////////

func (r *dataSource) Create(ctx context.Context, key string, value any) error {
	return r.saveJSON(ctx, key, value, 0)
}

func (r *dataSource) Update(ctx context.Context, key string, value any) error {
	return r.saveJSON(ctx, key, value, 0)
}

func (r *dataSource) Delete(ctx context.Context, key string) error {
	return r.RDb.Del(ctx, key).Err()
}

func (r *dataSource) Get(ctx context.Context, key string, dest any) error {
	b, err := r.RDb.Get(ctx, key).Bytes()
	if err != nil {
		return err // includes redis.Nil when not found
	}
	return json.Unmarshal(b, dest)
}

////////////////////////////////////////////////////////////
// -------------------- Save Internal --------------------
////////////////////////////////////////////////////////////

func (r *dataSource) saveJSON(
	ctx context.Context,
	key string,
	value any,
	ttl time.Duration,
) error {
	if value == nil {
		return errors.New("nil value")
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// SET handles TTL directly; no need for separate EXPIRE unless you want 2-step
	return r.RDb.Set(ctx, key, b, ttl).Err()
}

////////////////////////////////////////////////////////////
// -------------------- Index Ops ------------------------
////////////////////////////////////////////////////////////

func (r *dataSource) AddToIndex(ctx context.Context, indexKey, id string, score float64) error {
	return r.RDb.ZAdd(ctx, indexKey, redis.Z{
		Score:  score,
		Member: id,
	}).Err()
}

func (r *dataSource) RemoveFromIndex(ctx context.Context, indexKey, id string) error {
	return r.RDb.ZRem(ctx, indexKey, id).Err()
}

func (r *dataSource) GetAllFromIndex(
	ctx context.Context,
	indexKey string,
	start, stop int64,
) ([]string, error) {
	return r.RDb.ZRevRange(ctx, indexKey, start, stop).Result()
}

////////////////////////////////////////////////////////////
// -------------------- Atomic Helpers -------------------
////////////////////////////////////////////////////////////

func (r *dataSource) UpsertWithIndex(
	ctx context.Context,
	key, indexKey, id string,
	score float64,
	value any,
	ttl time.Duration,
) error {
	if value == nil {
		return errors.New("nil value")
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	pipe := r.RDb.TxPipeline()
	pipe.Set(ctx, key, b, ttl)
	pipe.ZAdd(ctx, indexKey, redis.Z{
		Score:  score,
		Member: id,
	})
	_, err = pipe.Exec(ctx)
	return err
}

func (r *dataSource) DeleteWithIndex(
	ctx context.Context,
	key, indexKey, id string,
) error {
	pipe := r.RDb.TxPipeline()
	pipe.Del(ctx, key)
	pipe.ZRem(ctx, indexKey, id)
	_, err := pipe.Exec(ctx)
	return err
}

////////////////////////////////////////////////////////////
// -------------------- Bulk Read ------------------------
////////////////////////////////////////////////////////////

func (r *dataSource) GetAll(
	ctx context.Context,
	indexKey string,
	destSlice any,
) error {
	ids, err := r.RDb.ZRevRange(ctx, indexKey, 0, -1).Result()
	if err != nil {
		return err
	}

	slicePtr := reflect.ValueOf(destSlice)
	if slicePtr.Kind() != reflect.Ptr || slicePtr.Elem().Kind() != reflect.Slice {
		return errors.New("destSlice must be pointer to slice")
	}

	elemType := slicePtr.Elem().Type().Elem()
	result := reflect.MakeSlice(slicePtr.Elem().Type(), 0, len(ids))

	if len(ids) == 0 {
		slicePtr.Elem().Set(result)
		return nil
	}

	// MGET faster than piping GET in a loop
	vals, err := r.RDb.MGet(ctx, ids...).Result()
	if err != nil {
		return err
	}

	// ids that no longer have values (expired/deleted)
	var toRemove []any

	for i, id := range ids {
		v := vals[i]
		if v == nil {
			toRemove = append(toRemove, id)
			continue
		}

		var raw []byte
		switch t := v.(type) {
		case string:
			raw = []byte(t)
		case []byte:
			raw = t
		default:
			// go-redis typically returns string; but handle defensively
			b, e := json.Marshal(t)
			if e != nil {
				return e
			}
			raw = b
		}

		elemPtr := reflect.New(elemType)
		if err := json.Unmarshal(raw, elemPtr.Interface()); err != nil {
			return err
		}
		result = reflect.Append(result, elemPtr.Elem())
	}

	if len(toRemove) > 0 {
		_ = r.RDb.ZRem(ctx, indexKey, toRemove...).Err()
	}

	slicePtr.Elem().Set(result)
	return nil
}
