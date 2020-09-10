package utils

import (
	"crypto/rand"
)

// GenerateRandomNumber : generate random number
func GenerateRandomNumber() string {
	number, _ := rand.Prime(rand.Reader, 64)
	return number.String()
}
