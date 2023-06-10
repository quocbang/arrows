package server

import (
	"context"

	"github.com/quocbang/arrows/pkg/protobuf/authentication"
	"github.com/rs/xid"
)

func (s Server) Login(context.Context, *authentication.LoginRequest) (*authentication.LoginReply, error) {
	return &authentication.LoginReply{
		APIKey: xid.New().String(),
	}, nil
}
