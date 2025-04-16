package server

import (
	"net/http"
	"strings"

	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/UnLess24/coin/server/internal/currency"
	"github.com/gin-gonic/gin"
)

func currencyHandler(cdb cache.Cache) func(*gin.Context) {
	return func(c *gin.Context) {
		param := currency.CurrencyName(strings.ToUpper(c.Param("currency")))

		list, err := cache.List(c.Request.Context(), cdb)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		items := make([]currency.CurrencyValue, 0, len(list.Data))
		for _, item := range list.Data {
			if val, ok := item.Quote[param]; ok {
				items = append(items, val)
			}
		}

		c.JSON(http.StatusOK, items)
	}
}
