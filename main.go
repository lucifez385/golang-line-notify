package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lucifez385/golang-line-notify/router"
)

func main() {
	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Server is running 👋!")
	})

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "8000"
	}

	router.V1(app)
	// Start server on http://${heroku-url}:${port}
	log.Fatal(app.Listen(":" + port))

}
