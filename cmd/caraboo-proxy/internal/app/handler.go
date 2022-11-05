package app

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/liampulles/juryrig/cmd/caraboo-proxy/internal/store"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	cfg   Config
	store store.Store
}

func NewHandler(
	cfg Config,
	store *store.InMemoryStore,
) *Handler {
	return &Handler{
		cfg:   cfg,
		store: store,
	}
}

func (h *Handler) handle(c *fiber.Ctx) error {
	forwardURL, err := url.JoinPath(h.cfg.ForwardBaseURL, c.Path())
	if err != nil {
		log.Err(err).Msg("could not create forward url")
		return err
	}

	// Try forward the request
	err = proxy.Do(c, forwardURL)
	if err == nil && isFailStatus(c.Response().StatusCode()) {
		err = fmt.Errorf("bad status: %d", c.Response().StatusCode())
	}

	if err != nil {
		log.Debug().Str("err", err.Error()).Msg("proxy failed, using default")

		// It didn't work - let's try and load a previous
		ok := h.store.Get(c.Request(), c.Response())
		if ok {
			return nil
		}

		// Well lets send a fail message back then
		return c.Status(http.StatusBadGateway).SendString("oops - the backend is unavailable and we have nothing saved for this...\n")
	}

	// The forward worked! Record the result for later...
	h.store.Set(c.Request(), c.Response())

	return nil
}

func isFailStatus(status int) bool {
	return status == 502
}
