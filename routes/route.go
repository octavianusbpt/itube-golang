package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/octavianusbpt/itube-golang/controllers"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.GetUser)
}
