package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/currency"
	"github.com/redis/go-redis/v9"
)

func Update(ctx context.Context, cache *redis.Client, cfg *config.Config) {
	ticker := time.NewTicker(cfg.TickerDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ctxTimeout, cancel := context.WithTimeout(ctx, cfg.RequestTimeout)
			list, err := currency.LatestList(ctxTimeout, cfg)
			if err != nil {
				cancel()
				slog.Error("can`t get list data", "ERROR", err)
				continue
			}
			err = cache.Set(ctx, "list", list, 0).Err()
			if err != nil {
				slog.Error("can`t set list data to cache", "ERROR", err)
			}
			cancel()
		}
	}
}
