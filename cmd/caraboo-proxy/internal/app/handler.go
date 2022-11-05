package app

import "github.com/gofiber/fiber/v2"

func handler(c *fiber.Ctx) error {
	return c.Status(200).SendString("Hi there!")
}
