package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/store"
	"github.com/rs/zerolog/log"
)

func Run(cfg Config) error {
	app := fiber.New(fiber.Config{
		StrictRouting:         true,
		CaseSensitive:         true,
		AppName:               "caraboo-proxy",
		DisableStartupMessage: true,
	})
	store := store.NewInMemoryStore()
	handler := NewHandler(cfg, store)

	// Match any route
	app.Use(handler.handle)

	log.Info().Msgf("listening on port %d", cfg.Port)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Err(err).Msg("could not start fiber")
		return err
	}
	return nil
}

func LoadConfig(source goConfig.Source) (Config, error) {
	// Define defaults
	cfg := Config{
		Port:           9080,
		ForwardBaseURL: "http://127.0.0.1:80",
	}

	// Read from source
	typedSource := goConfig.NewTypedSource(source)
	if err := goConfig.LoadProperties(typedSource,
		goConfig.IntProp("PORT", &cfg.Port, false),
		goConfig.StrProp("FORWARD_BASE_URL", &cfg.ForwardBaseURL, false),
	); err != nil {
		log.Err(err).Msg("could not load config")
		return Config{}, err
	}

	return cfg, nil
}

type Config struct {
	Port           int
	ForwardBaseURL string
}
