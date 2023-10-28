package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordHelper(value string) string {
	generatedPassword, generatedPasswordError := bcrypt.GenerateFromPassword([]byte(value), 14)
	if generatedPasswordError != nil {
		log.Fatal("Error while hashing password", generatedPassword)
	}

	return string(generatedPassword)
}

func VerifyPassword(password string, enteredPassword string) bool {
	passwordError := bcrypt.CompareHashAndPassword([]byte(password), []byte(enteredPassword))
	return passwordError == nil
}
