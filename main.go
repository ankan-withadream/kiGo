package main

import (
	// Std library to handel http request on Go
	"github.com/gofiber/fiber/v3"
)

const PORT = ":8080"

func main() {

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// http request listener
	app.Listen(PORT)
}
