package cache

import "context"

const (
	LatestListKey = "latest:list"
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
}
