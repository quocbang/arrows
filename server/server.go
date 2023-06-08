package server

import (
	"context"

	"github.com/quocbang/arrows/pkg/protobuf/authentication"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type Server struct {
	dm        string
	getLogger func(context.Context) *zap.Logger
}

type ServerParams struct {
	LocalDataManager string
	Logger           func(context.Context) *zap.Logger
}

func NewServer(params ServerParams) *Server {
	return &Server{
		dm:        params.LocalDataManager,
		getLogger: params.Logger,
	}
}

func (s *Server) Login(context.Context, *authentication.LoginRequest) (*authentication.LoginReply, error) {
	return &authentication.LoginReply{
		APIKey: xid.New().String(),
	}, nil
}
