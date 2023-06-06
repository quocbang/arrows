package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"

	"github.com/quocbang/arrows/config"
)

func main() {
	flags := parseFlag()
	fmt.Println(flags)
}

type FlagConfig struct {
	Options    config.FlagOptions
	TLSOptions config.TLSOptionsType
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
