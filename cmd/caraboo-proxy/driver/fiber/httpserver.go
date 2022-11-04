package fiber

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	port int
}

func NewHTTPServer(port int) *HTTPServer {
	return &HTTPServer{
		port: port,
	}
}

func (server *HTTPServer) Run() error {
	app := fiber.New()

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("First handler")
		return c.Next()
	})

	return app.Listen(fmt.Sprintf(":%d", server.port))
}
