package helpers

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func stringWithCharset(length int, charset string) string {
	rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, length)

	for index := range bytes {
		bytes[index] = charset[rand.Intn(len(charset))]
	}

	return string(bytes)
}

func Generate16CharUUID() string {
	return stringWithCharset(16, charset)
}
