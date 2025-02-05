package cache

import (
	"context"
	"fmt"

	"github.com/UnLess24/coin/server/config"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedis(cfg *config.Config) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.DB,       // use default DB
	})
	return &RedisCache{client: client}
}

func (r *RedisCache) Set(ctx context.Context, key string, value []byte) error {
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("can't get data from redis by key: %v %w", key, err)
	}
	return data, nil
}
