package server

import (
	"log/slog"
	"net"

	"github.com/UnLess24/coin/server/internal/cache"
	"github.com/UnLess24/coin/server/internal/services"
	"github.com/UnLess24/coin/server/pkg/api/coin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	srv *grpc.Server
	lis net.Listener
}

func (s *grpcServer) Serve() error {
	return s.srv.Serve(s.lis)
}

func (s *grpcServer) Shutdown() error {
	s.srv.GracefulStop()
	return nil
}

func (s *grpcServer) Close() error {
	return s.lis.Close()
}

func NewGRPC(c cache.Cache, addr string) *grpcServer {
	srv := grpc.NewServer()
	service := services.NewCoinService(c)
	coin.RegisterCoinServiceServer(srv, service)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("exit reason", "ERROR", err)
	}

	reflection.Register(srv)

	slog.Info("grpc server started", "port", addr)

	return &grpcServer{srv: srv, lis: lis}
}
