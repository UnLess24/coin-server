package server

import (
	"net/http"

	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/gin-gonic/gin"
)

func listHandler(cdb cache.Cache) func(*gin.Context) {
	return func(c *gin.Context) {
		list, err := cache.List(c.Request.Context(), cdb)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, list)
	}
}
