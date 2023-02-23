package helpers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogIfError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

func ResponseIfError(err error, errorCode int, errorMessage string) {
	if err != nil {
		fiber.NewError(errorCode, errorMessage)
		return
	}
}
