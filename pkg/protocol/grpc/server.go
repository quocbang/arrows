package grpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/quocbang/arrows/config"
	api "github.com/quocbang/arrows/pkg/protobuf/api"
	"github.com/quocbang/arrows/server"
)

const (
	ArrowsAPIToken = "arrows-api-token"
)

type GRPCServer struct {
	GRPCHost string
	GRPCPort int
	TLS      config.TLSOptionsType
	Config   config.Config
}

func (t *GRPCServer) Run() error {
	defer zap.L().Sync()

	opts := []grpc.ServerOption{}

	tcpGRPCServer := grpc.NewServer(opts...)
	activeServer := server.NewServer(server.ServerParams{
		LocalDataManager: "a",
		Logger: func(ctx context.Context) *zap.Logger {
			return &zap.Logger{}
		},
	})

	// Register API
	api.RegisterAPIServer(tcpGRPCServer, activeServer)

	address := fmt.Sprintf("%s:%d", t.GRPCHost, t.GRPCPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	zap.L().Info("serving gRPC server..", zap.String("address", address))

	return tcpGRPCServer.Serve(lis)
}
