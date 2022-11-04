package wire

import (
	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/driver/fiber"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/usecase"
	"github.com/rs/zerolog/log"
)

type App struct {
	httpserver *fiber.HTTPServer
}

func (app *App) Run() error {
	return app.httpserver.Run()
}

func Wire(cfg Config) *App {
	svc := usecase.NewService()

	httpserver := fiber.NewHTTPServer(cfg.Port, svc)

	return &App{
		httpserver: httpserver,
	}
}

func LoadConfig(source goConfig.Source) (Config, error) {
	// Define defaults
	cfg := Config{
		Port: 9080,
	}

	// Read from source
	typedSource := goConfig.NewTypedSource(source)
	if err := goConfig.LoadProperties(typedSource,
		goConfig.IntProp("PORT", &cfg.Port, false),
	); err != nil {
		log.Err(err).Msg("could not load config")
		return Config{}, err
	}

	return cfg, nil
}

type Config struct {
	Port int
}
