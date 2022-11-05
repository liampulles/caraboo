package app

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handle(c *fiber.Ctx) error {
	return c.Status(200).SendString("Hi there!")
}
