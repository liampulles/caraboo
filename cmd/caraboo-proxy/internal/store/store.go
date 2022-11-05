package store

import "github.com/gofiber/fiber/v2"

type Store interface {
	Set(request *fiber.Request, copyFrom *fiber.Response)
	Get(request *fiber.Request, into *fiber.Response) bool
}
