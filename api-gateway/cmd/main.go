package main

import (
	"gateway/gateway"
	"gateway/server"
	"gateway/util"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configPath, found := os.LookupEnv("CONFIG_PATH")

	if !found {
		configPath = "."
	}

	config, err := util.LoadConfig(configPath)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	go gateway.RunGatewayServer(config)
	server.RunGrpcServer(config)
}
