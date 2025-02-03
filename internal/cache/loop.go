package cache

import (
	"context"
	"log/slog"
	"time"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/currency"
)

func Update(ctx context.Context, cfg *config.Config) {
	ticker := time.NewTicker(cfg.TickerDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			list, err := currency.SendRequest(ctx, cfg)
			if err != nil {
				slog.Error("can`t get list data", "ERROR", err)
				continue
			}
			slog.Info("list data", "ID", slog.IntValue(list.Data[0].Id))
		}
	}
}
