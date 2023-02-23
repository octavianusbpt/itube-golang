package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/octavianusbpt/itube-golang/controllers"
	"github.com/octavianusbpt/itube-golang/middleware"
)

func Routes(app *fiber.App) {
	app.Post("/signup", controllers.Signup)                            // Signup
	app.Post("/login", controllers.Login)                              // Login
	app.Get("/validate", middleware.RequireAuth, controllers.Validate) // Validate authorized access
}
