package server

import (
	"context"

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

func NewServer(params ServerParams) Server {
	return Server{
		dm:        params.LocalDataManager,
		getLogger: params.Logger,
	}
}
