package rest

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/quocbang/arrows/config"
	api "github.com/quocbang/arrows/pkg/protobuf/api"
)

type RestServer struct {
	Host               string
	Port               int
	GRPCServerEndpoint string
	TLS                config.TLSOptionsType
	Config             config.Config
}

func (s *RestServer) gatewayMux() (*runtime.ServeMux, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(h string) (string, bool) {
			return h, true
		}),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	)

	var opts []grpc.DialOption

	if s.TLS.UseTLS() {
		var cred credentials.TransportCredentials
		cred = credentials.NewTLS(&tls.Config{})

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(cred),
		}
	} else {
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
		}
	}

	// Register arrows handlers
	if err := api.RegisterAPIHandlerFromEndpoint(ctx, mux, s.GRPCServerEndpoint, opts); err != nil {
		return nil, cancel, err
	}

	return mux, cancel, nil
}

// Run serves RESTful server.
func (s *RestServer) Run() error {
	defer zap.L().Sync() // nolint

	mux, cancel, err := s.gatewayMux()
	if err != nil {
		return err
	}
	defer cancel()

	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	zap.L().Info("serving gateway server..", zap.String("address", address))

	if s.TLS.UseTLS() {
		return srv.ListenAndServeTLS(s.TLS.Cert, s.TLS.Key)
	}

	return srv.ListenAndServe()
}
