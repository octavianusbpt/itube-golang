package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/octavianusbpt/itube-golang/database"
	"github.com/octavianusbpt/itube-golang/helpers"
	"github.com/octavianusbpt/itube-golang/models"
)

func RequireAuth(c *fiber.Ctx) error {
	// Attempt to retrieve cookie
	tokenString := c.Cookies("Authorization")

	if tokenString == "" {
		return fiber.ErrUnauthorized
	}

	// Parse and validate cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})
	helpers.PanicIfError(err)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if cookie token is expired
		if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64) {
			return fiber.ErrUnauthorized
		}

		// Find the corresponding user
		var user models.User
		database.DB.First(&user, claims["ID"])

		if user.ID == 0 {
			return fiber.ErrNotFound
		}
	} else {
		return fiber.ErrUnauthorized
	}

	return c.Next()

}
