package cache

import (
	"context"
	"fmt"

	"github.com/UnLess24/coin/server/internal/currency"
)

const (
	LatestListKey = "latest:list"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
}

// List get list from cache
func List(ctx context.Context, c Cache) (currency.ListResponse, error) {
	data, err := c.Get(ctx, LatestListKey)
	if err != nil {
		return currency.ListResponse{}, fmt.Errorf("failed to get data from cache %w", err)
	}
	list, err := currency.MakeList(data)
	if err != nil {
		return currency.ListResponse{}, fmt.Errorf("failed to make list %w", err)
	}
	return list, nil
}
