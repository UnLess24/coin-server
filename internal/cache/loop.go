package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/currency"
)

func Update(ctx context.Context, c Cache, cfg *config.Config) {
	ticker := time.NewTicker(cfg.TickerDuration)
	defer ticker.Stop()

	makeRequest(ctx, c, cfg)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			makeRequest(ctx, c, cfg)
		}
	}
}

func makeRequest(ctx context.Context, c Cache, cfg *config.Config) {
	ctxTimeout, cancel := context.WithTimeout(ctx, cfg.RequestTimeout)
	defer cancel()

	list, err := currency.LatestList(ctxTimeout, cfg)
	if err != nil {
		slog.Error("can`t get list data", "ERROR", err)
		return
	}

	err = c.Set(ctx, LatestListKey, list)
	if err != nil {
		slog.Error("can`t set list data to cache", "ERROR", err)
	}
}
