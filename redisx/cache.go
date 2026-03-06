package redisx

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string, value interface{}) error
	GetCache(ctx context.Context, key string, value interface{}, ttl *time.Duration, loadFunc func() (interface{}, error)) error
	Set(ctx context.Context, key string, cacheItem *Item) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) bool
	IsErrCacheMiss(err error) bool
	ListKeysByPrefix(ctx context.Context, prefix string) ([]string, error)
	DeleteByPrefix(ctx context.Context, prefix string) error
}

type Item struct {
	Value interface{}
	TTL   *time.Duration
}

type cacheX struct {
	currentCache *cache.Cache
	client       redis.UniversalClient
}

func NewCache(client redis.UniversalClient) Cache {
	c := cache.New(&cache.Options{
		Redis:      client,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
		Marshal:    json.Marshal,
		Unmarshal:  json.Unmarshal,
	})
	return &cacheX{currentCache: c, client: client}
}

func (c *cacheX) GetCache(ctx context.Context, key string, value interface{}, ttl *time.Duration, loadFunc func() (interface{}, error)) error {
	item := &cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		Do: func(i *cache.Item) (interface{}, error) {
			return loadFunc()
		},
	}
	if ttl != nil {
		item.TTL = *ttl
	}

	err := c.currentCache.Once(item)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheX) Get(ctx context.Context, key string, value interface{}) error {
	err := c.currentCache.Get(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheX) Set(ctx context.Context, key string, cacheItem *Item) error {
	item := &cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: cacheItem.Value,
	}
	if cacheItem.TTL != nil {
		item.TTL = *cacheItem.TTL
	}
	err := c.currentCache.Set(item)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheX) Delete(ctx context.Context, key string) error {
	err := c.currentCache.Delete(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

func (c *cacheX) Exists(ctx context.Context, key string) bool {
	return c.currentCache.Exists(ctx, key)
}

func (c *cacheX) IsErrCacheMiss(err error) bool {
	return errors.Is(err, cache.ErrCacheMiss)
}

func (c *cacheX) ListKeysByPrefix(ctx context.Context, prefix string) ([]string, error) {
	iter := c.client.Scan(ctx, 0, prefix+":*", 0).Iterator()
	var keys []string
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return keys, nil
}

func (c *cacheX) DeleteByPrefix(ctx context.Context, prefix string) error {
	iter := c.client.Scan(ctx, 0, prefix+":*", 0).Iterator()
	for iter.Next(ctx) {
		err := c.Delete(ctx, iter.Val())
		if err != nil {
			if c.IsErrCacheMiss(err) {
				continue
			}
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}
