package router

import (
	"github.com/lucifez385/golang-line-notify/controller"

	"github.com/gofiber/fiber/v2"
)

func V1(app *fiber.App) {
	v1 := app.Group("api/v1")
	lineNotify := v1.Group("/line-notify")
	lineNotifyController := controller.NewLineNotify()
	lineNotify.Post("/:token", lineNotifyController.SendNotify)
}
