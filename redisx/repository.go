package redisx

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Repository[T any] struct {
	ds      DataSource
	name    string
	ttl     time.Duration
	env     string
	getIDFn func(*T) string
}

func NewRepository[T any](
	ds DataSource,
	name string,
	env string,
	getIDFn func(*T) string,
) *Repository[T] {
	return &Repository[T]{
		ds:      ds,
		name:    name,
		env:     env,
		getIDFn: getIDFn,
	}
}

func (s *Repository[T]) WithTTL(ttl time.Duration) *Repository[T] {
	s.ttl = ttl
	return s
}

func (s *Repository[T]) recordKey(id string) string {
	return fmt.Sprintf("%s:%s:%s", s.name, s.env, id)
}

func (s *Repository[T]) indexKey() string {
	return fmt.Sprintf("%s:%s:index", s.name, s.env)
}

func (s *Repository[T]) Create(ctx context.Context, obj *T) error {
	id := s.getIDFn(obj)
	key := s.recordKey(id)

	return s.ds.UpsertWithIndex(
		ctx,
		key,
		s.indexKey(),
		key,
		float64(time.Now().Unix()),
		obj,
		s.ttl,
	)
}

func (s *Repository[T]) CreateList(ctx context.Context, obj *[]T) error {
	if obj == nil {
		return errors.New("object is null")
	}

	var errs []error

	for i := range *obj {
		if err := s.Create(ctx, &(*obj)[i]); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (s *Repository[T]) Update(ctx context.Context, obj *T) error {
	return s.Create(ctx, obj)
}

func (s *Repository[T]) Delete(ctx context.Context, id string) error {
	key := s.recordKey(id)
	return s.ds.DeleteWithIndex(ctx, key, s.indexKey(), key)
}

func (s *Repository[T]) FindOne(ctx context.Context, id string) (*T, error) {
	key := s.recordKey(id)
	result := new(T)

	err := s.ds.Get(ctx, key, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Repository[T]) FindList(ctx context.Context) ([]T, error) {
	var result []T
	err := s.ds.GetAll(ctx, s.indexKey(), &result)
	return result, err
}
