package cache

import (
	"github.com/UnLess24/coin/server/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.DB,       // use default DB
	})
}
