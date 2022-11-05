package app

import (
	"os"

	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg, err := app.LoadConfig(goConfig.NewEnvSource())
	if err != nil {
		os.Exit(1)
	}

	if err := app.Run(cfg); err != nil {
		os.Exit(2)
	}
}
