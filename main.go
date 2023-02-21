package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/octavianusbpt/itube-golang/database"
	"github.com/octavianusbpt/itube-golang/helpers"
	"github.com/octavianusbpt/itube-golang/routes"
)

func init() {

	// Initialize environment variables
	err := godotenv.Load()
	helpers.LogIfError(err, "Error loading .env file")

	// Initialize database
	database.InitializeDatabase()
	database.SyncDB()
}

func main() {

	// Setup App
	app := fiber.New()

	// Configure public
	app.Static("/", "./public")

	// Routing
	routes.Routes(app)

	// Start App
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
