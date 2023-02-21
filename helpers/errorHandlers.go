package helpers

import "log"

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
