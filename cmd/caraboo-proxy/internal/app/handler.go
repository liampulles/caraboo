package app

import "github.com/gofiber/fiber/v2"

type Handler struct {
	cfg Config
}

func NewHandler(cfg Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) handle(c *fiber.Ctx) error {
	return c.Status(200).SendString("Hi there!")
}
