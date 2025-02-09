package server

import (
	"net/http"

	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/gin-gonic/gin"
)

func New(c cache.Cache, addr string) *http.Server {
	r := gin.Default()

	r.GET("/list", listHandler(c))

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
