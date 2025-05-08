package server

import (
	"context"
	"net/http"

	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/UnLess24/coin/server/internal/handler"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	srv *http.Server
}

func (s *httpServer) Serve() error {
	return s.srv.ListenAndServe()
}

func (s *httpServer) Shutdown() error {
	return s.srv.Shutdown(context.Background())
}

func (s *httpServer) Close() error {
	return s.srv.Close()
}

func NewHTTP(c cache.Cache, addr string) *httpServer {
	r := gin.Default()

	r.GET("/list", handler.List(c))
	r.GET("/currency", handler.Currency(c))

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &httpServer{srv: srv}
}
