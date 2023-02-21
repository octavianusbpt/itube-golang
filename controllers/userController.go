package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) {
	// var body struct {
	// 	Name     string
	// 	Email    string
	// 	Password string
	// }
}

func GetUser(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
