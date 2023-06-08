package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/quocbang/arrows/config"
	tcpAPI "github.com/quocbang/arrows/pkg/protobuf/api/tcp"
	"github.com/quocbang/arrows/server"
)

const (
	ArrowsAPIToken = "arrows-api-token"
)

type GRPCServer struct {
	Host     string
	Port     int
	GRPCHost string
	GRPCPort int
	TLS      config.TLSOptionsType
	Config   config.Config
}

func (t *GRPCServer) Run() error {
	defer zap.L().Sync()

	opts := []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Minute * 2),
	}

	tcpGRPCServer := grpc.NewServer(opts...)

	// Register API
	tcpAPI.RegisterTCPArrowsServer(grpc.NewServer(opts...),
		server.NewServer(server.ServerParams{
			LocalDataManager: "a",
			Logger: func(ctx context.Context) *zap.Logger {
				return &zap.Logger{}
			},
		}))

	address := fmt.Sprintf("%s:%d", t.GRPCHost, t.GRPCPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		zap.L().Error("failed to listen: %v", zap.Error(err))
		return err
	}
	zap.L().Info("serving gRPC server..", zap.String("address", address))

	return tcpGRPCServer.Serve(lis)
}
