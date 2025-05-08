package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/UnLess24/coin/server/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.MustRead()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer cancel()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
	}()

	cdb := cache.NewRedis(cfg)
	go cache.Update(ctx, cdb, cfg)

	srv, err := server.New(cdb, cfg.Server.Port, cfg.Server.Type)
	if err != nil {
		slog.Error("exit reason", "ERROR", err)
		return
	}
	defer srv.Close()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return srv.Serve()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return srv.Shutdown()
	})

	if err := g.Wait(); err != nil {
		slog.Error("exit reason", "ERROR", err)
	}
}
