package wire

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	goConfig "github.com/liampulles/go-config"
	"github.com/rs/zerolog/log"
)

func Run(cfg Config) error {
	app := fiber.New(fiber.Config{
		StrictRouting:         true,
		CaseSensitive:         true,
		AppName:               "caraboo-proxy",
		DisableStartupMessage: true,
	})

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hi there!")
	})

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
