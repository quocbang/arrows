package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v3"

	"github.com/quocbang/arrows/config"
)

func main() {
	flags := parseFlag()

	config, err := loadConfig(flags.Options.Config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(flags)
	log.Println(*config.PostGres)
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
