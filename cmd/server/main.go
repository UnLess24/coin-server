package main

import (
	"context"
	"net/http"

	"github.com/UnLess24/coin/server/config"
	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.MustRead()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cdb := cache.NewRedis(cfg)
	go cache.Update(ctx, cdb, cfg)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		list, err := cache.List(ctx, cdb)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, list)
	})
	r.Run(cfg.Server.Port)
}
