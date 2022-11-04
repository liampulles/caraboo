package wire

import (
	"fmt"

	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/driver/fiber"
)

type App struct {
	httpserver *fiber.HTTPServer
}

func (app *App) Run() error {
	return app.httpserver.Run()
}

func Wire(cfg Config) *App {
	httpserver := fiber.NewHTTPServer(cfg.Port)

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
		return Config{}, fmt.Errorf("could not fetch config: %w", err)
	}

	return cfg, nil
}

type Config struct {
	Port int
}
