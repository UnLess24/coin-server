package server

import (
	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/gin-gonic/gin"
)

func New(c cache.Cache) *gin.Engine {
	r := gin.Default()

	r.GET("/list", listHandler(c))

	return r
}
