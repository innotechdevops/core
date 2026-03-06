package redisx

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/extra/redisotel/v9"

	"github.com/redis/go-redis/v9"
)

var RedisNil = redis.Nil

type Client interface {
	Publish(ctx context.Context, channel string, message interface{}) error
	Subscribe(ctx context.Context, channel string) *redis.PubSub
	PSubscribe(ctx context.Context, channel string) *redis.PubSub
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	SubscribeMessage(ctx context.Context, options *SubscribeMessageOptions) func()
	PSubscribeMessage(ctx context.Context, channel string, userIdPrefix string, onMessage func(id, message string)) func()
	Client() redis.UniversalClient
	Close() error
}

type SubscribeMessageOptions struct {
	Channel     string
	OnMessage   func(message string)
	ChannelSize int
}

type RedisOptions struct {
	ClusterMode bool
	Addr        string
	DB          int
	Password    string
	TraceEnable bool
}

type redisClient struct {
	client redis.UniversalClient
}

func (r *redisClient) Client() redis.UniversalClient {
	return r.client
}

func NewClient(options *RedisOptions) (Client, error) {
	var c redis.UniversalClient
	if options.ClusterMode {
		c = newClusterClient(options)
	} else {
		c = newClient(options)
	}

	err := c.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	if options.TraceEnable {
		err = redisotel.InstrumentTracing(c)
		if err != nil {
			return nil, err
		}
	}

	return &redisClient{client: c}, nil
}

func newClient(options *RedisOptions) redis.UniversalClient {
	redisOptions := &redis.Options{
		Addr:       options.Addr,
		DB:         options.DB,
		MaxRetries: 3,
	}

	if options.Password != "" {
		redisOptions.Password = options.Password
	}

	return redis.NewClient(redisOptions)
}

func newClusterClient(options *RedisOptions) redis.UniversalClient {
	redisOptions := &redis.ClusterOptions{
		Addrs: strings.Split(options.Addr, ","),
	}

	if options.Password != "" {
		redisOptions.Password = options.Password
	}

	return redis.NewClusterClient(redisOptions)
}

func (r *redisClient) Publish(ctx context.Context, channel string, message interface{}) error {
	return r.client.Publish(ctx, channel, message).Err()
}

func (r *redisClient) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.client.Subscribe(ctx, channel)
}

func (r *redisClient) PSubscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.client.PSubscribe(ctx, channel)
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisClient) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redisClient) SubscribeMessage(ctx context.Context, options *SubscribeMessageOptions) func() {
	pubSub := r.Subscribe(ctx, options.Channel)

	channelSize := 100
	if options.ChannelSize > 0 {
		channelSize = options.ChannelSize
	}

	ch := pubSub.Channel(redis.WithChannelSize(channelSize))

	go func() {
		for msg := range ch {
			options.OnMessage(msg.Payload)
		}
	}()

	return func() {
		_ = pubSub.Close()
	}
}

func (r *redisClient) PSubscribeMessage(ctx context.Context, channel string, userIdPrefix string, onMessage func(id, message string)) func() {
	pubSub := r.PSubscribe(ctx, channel+userIdPrefix)
	ch := pubSub.Channel()
	go func() {
		for msg := range ch {
			idx := strings.Index(msg.Channel, userIdPrefix)
			id := msg.Channel[idx+len(userIdPrefix):]
			onMessage(id, msg.Payload)
		}
	}()

	return func() {
		_ = pubSub.Close()
	}
}

func (r *redisClient) Close() error {
	return r.client.Close()
}
