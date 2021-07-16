package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucifez385/golang-line-notify/router"
)

func main() {
	app := fiber.New()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("Server is running 👋!")
	})

	router.V1(app)
	app.Listen(":3000")
	fmt.Println("Server is running")
}
