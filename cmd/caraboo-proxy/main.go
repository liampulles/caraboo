package main

import (
	"os"

	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := wire.LoadConfig(goConfig.NewEnvSource())
	if err != nil {
		os.Exit(1)
	}

	if err := wire.Run(cfg); err != nil {
		os.Exit(2)
	}
}
