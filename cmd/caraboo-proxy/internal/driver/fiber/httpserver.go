package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/usecase"
	"github.com/rs/zerolog/log"
)

type HTTPServer struct {
	port int
	svc  *usecase.Service
}

func NewHTTPServer(
	port int,
	svc *usecase.Service,
) *HTTPServer {
	return &HTTPServer{
		port: port,
		svc:  svc,
	}
}

func (server *HTTPServer) Run() error {
	app := fiber.New(fiber.Config{
		StrictRouting:         true,
		CaseSensitive:         true,
		AppName:               "caraboo-proxy",
		DisableStartupMessage: true,
	})

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		req := &usecase.Request{
			Path: c.Path(),
		}

		res := server.svc.ProcessRequest(req)

		return c.Status(res.Status).Send(res.Body)
	})

	log.Info().Msgf("listening on port %d", server.port)
	if err := app.Listen(fmt.Sprintf(":%d", server.port)); err != nil {
		log.Err(err).Msg("could not start fiber")
		return err
	}
	return nil
}
