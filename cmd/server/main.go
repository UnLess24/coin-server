package main

import (
	"context"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/UnLess24/coin/server/internal/server"
)

func main() {
	cfg := config.MustRead()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cdb := cache.NewRedis(cfg)
	go cache.Update(ctx, cdb, cfg)

	r := server.New(cdb)
	if err := r.Run(cfg.Server.Port); err != nil {
		panic(err)
	}
}
