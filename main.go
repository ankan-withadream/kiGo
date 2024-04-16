package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
)

const PORT = ":8080"

type Response struct {
}

func main() {

	app := fiber.New()

	app.Get("/kigo", func(c fiber.Ctx) error {

		response, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=43272343095c48fb98d7463190d2cf2f")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Print(err.Error())
		}

		// return c.SendString("Hello, World!")
		return c.Send(bodyBytes)
	})

	// http request listener
	app.Listen(PORT)
}
