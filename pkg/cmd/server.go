package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/quocbang/arrows/config"
	"github.com/quocbang/arrows/pkg/logger"
	"github.com/quocbang/arrows/pkg/protocol/tcp/grpc"
	"github.com/quocbang/arrows/pkg/protocol/tcp/rest"
)

func RunServer() {
	flags := parseFlag()

	// initialize logger.
	logger.Init(false)

	config, err := loadConfig(flags.Options.Config)
	if err != nil {
		log.Fatal(err)
	}

	if !UseTLS(flags.TLSOptions) {
		zap.L().Warn("serving service without TLS handshake")
	}

	wg := sync.WaitGroup{}
	defer wg.Wait()

	// listen with restful api.
	wg.Add(1)
	go func() {
		defer wg.Done()

		rest := rest.RestServer{
			Host:     flags.Options.Host,
			Port:     flags.Options.Port,
			GRPCHost: flags.Options.GRPCHost,
			GRPCPort: flags.Options.GRPCPort,
			TLS:      flags.TLSOptions,
			Config:   *config,
		}

		if err := rest.Run(); err != nil {
			zap.L().Error("failed to serve gRPC: ", zap.Error(err))
		}
		zap.L().Info("gRPC server stopped")
	}()

	// listen with grpc api.
	wg.Add(1)
	go func() {
		defer wg.Done()

		grpc := grpc.GRPCServer{
			Host:     flags.Options.Host,
			Port:     flags.Options.Port,
			GRPCHost: flags.Options.GRPCHost,
			GRPCPort: flags.Options.GRPCPort,
			TLS:      flags.TLSOptions,
			Config:   *config,
		}

		if err := grpc.Run(); err != nil {
			zap.L().Error("failed to serve gRPC: ", zap.Error(err))
		}
		zap.L().Error("gRPC server stopped")
	}()
}

type FlagConfig struct {
	Options    config.FlagOptions
	TLSOptions config.TLSOptionsType
}

func loadConfig(configPath string) (*config.Config, error) {
	conf := config.Config{}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return &config.Config{}, err
	}

	if err := yaml.Unmarshal(data, &conf); err != nil {
		return &config.Config{}, err
	}

	return &conf, nil
}

func parseFlag() *FlagConfig {
	var conf FlagConfig
	// set variable.
	configurations := []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Server Configuration",
			LongDescription:  "Server Configuration",
			Options:          &conf.Options,
		},
		{
			ShortDescription: "Server Configuration",
			LongDescription:  "Server Configuration",
			Options:          &conf.TLSOptions,
		},
	}

	// parse command line flags.
	parser := flags.NewParser(nil, flags.Default)
	for _, optsGroup := range configurations {
		if _, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options); err != nil {
			log.Fatal(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok && fe.Type == flags.ErrHelp {
			code = 0
		}
		os.Exit(code)
	}

	return &conf
}

func UseTLS(tls config.TLSOptionsType) bool {
	return tls.Key != "" && tls.Cert != ""
}
