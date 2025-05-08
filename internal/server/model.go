package server

import (
	"fmt"

	"github.com/UnLess24/coin/server/internal/cache"
)

const (
	GRPC = "grpc"
	HTTP = "http"
)

type Server interface {
	Serve() error
	Shutdown() error
	Close() error
}

func New(c cache.Cache, addr, serverType string) (Server, error) {
	switch serverType {
	case GRPC:
		return NewGRPC(c, addr), nil
	case HTTP:
		return NewHTTP(c, addr), nil
	default:
		return nil, fmt.Errorf("unknown server type %s", serverType)
	}
}
