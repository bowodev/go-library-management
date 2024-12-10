package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bowodev/go-library-management/config"
	"github.com/bowodev/go-library-management/internal/interfaces"
	"github.com/redis/go-redis/v9"
)

type redisClient[T any] struct {
	client *redis.Client
	key    string
	ttl    time.Duration
}

var _ interfaces.ICache[any] = (*redisClient[any])(nil)

func New[T any](cfg config.Config, key string, ttl time.Duration) (*redisClient[T], *redis.Client) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: "",
		DB:       cfg.Redis.DB,
	})

	return &redisClient[T]{
		client: rdb,
		key:    key,
		ttl:    ttl,
	}, rdb
}

func (r *redisClient[T]) Set(ctx context.Context, key string, value T) error {
	keyFormat := fmt.Sprintf("%s.%s", r.key, key)

	exists, _ := r.client.Get(ctx, keyFormat).Result()
	if exists == "" {
		return nil
	}

	return r.client.SetNX(ctx, r.key, value, r.ttl).Err()
}

func (r *redisClient[T]) Get(ctx context.Context, key string) (T, error) {
	var result T

	keyFormat := fmt.Sprintf("%s.%s", r.key, key)
	value, err := r.client.Get(ctx, keyFormat).Result()
	if err != nil {
		return result, err
	}
	if value == "" {
		return result, nil
	}

	return result, json.Unmarshal([]byte(value), &result)
}

func (r *redisClient[T]) Del(ctx context.Context, key string) error {
	keyFormat := fmt.Sprintf("%s.%s", r.key, key)

	return r.client.Del(ctx, keyFormat).Err()
}
