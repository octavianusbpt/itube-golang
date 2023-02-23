package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/octavianusbpt/itube-golang/database"
	"github.com/octavianusbpt/itube-golang/helpers"
	"github.com/octavianusbpt/itube-golang/models"
	"github.com/octavianusbpt/itube-golang/models/web"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {

	// Get data from request body
	body := new(web.UserSignup)

	err := c.BodyParser(&body)
	helpers.ResponseIfError(err, fiber.StatusBadRequest, "Failed to read body")

	// Password hashing using bcrypt (golang crypto package)
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	helpers.ResponseIfError(err, fiber.StatusNotImplemented, "Failed to hash password")

	// User Creation
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hashedPwd)}
	result := database.DB.Create(&user)

	helpers.ResponseIfError(result.Error, fiber.StatusNotImplemented, "Failed to create user")

	// Response to user
	return c.JSON(&web.ResponseBody{Code: fiber.StatusOK, Description: "Sign up succesful"})
}

func Login(c *fiber.Ctx) error {

	// Get data from request body
	body := new(web.UserLogin)

	err := c.BodyParser(&body)
	helpers.ResponseIfError(err, fiber.StatusBadRequest, "Failed to read body")

	// Identify the user by email
	var user models.User
	database.DB.Where(&models.User{Email: body.Email}).First(&user)

	if user.ID == 0 {
		return fiber.ErrNotFound
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	helpers.ResponseIfError(err, fiber.StatusUnauthorized, "Invalid credentials")

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ExpiresAt": time.Now().AddDate(0, 1, 0).Unix(),
		"Issuer":    "ITube",
		"IssuedAt":  time.Now().Unix(),
		"Subject":   "Login token",
		"ID":        user.ID,
	})

	ss, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
	helpers.ResponseIfError(err, fiber.StatusNotImplemented, "Failed to generate token")

	// Saves the cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = ss
	cookie.Expires = time.Now().AddDate(0, 1, 0)

	c.Cookie(cookie)

	// Send the response if login is successful
	return c.JSON(&web.ResponseBody{Code: fiber.StatusOK, Description: "Login succesful"})
}

func Validate(c *fiber.Ctx) error {
	return c.JSON(&web.ResponseBody{Code: fiber.StatusOK, Description: "Access granted"})
}
