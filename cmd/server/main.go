package main

import (
	"context"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/cache"
)

func main() {
	cfg := config.MustRead()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cache.Update(ctx, cfg)
}
