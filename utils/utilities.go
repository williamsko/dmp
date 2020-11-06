package utils

import (
	"crypto/rand"

	"github.com/gin-gonic/gin"
)

// GenerateRandomNumber : generate random number
func GenerateRandomNumber() string {
	number, _ := rand.Prime(rand.Reader, 64)
	return number.String()
}

// RespondWithError : respond api with error
func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}


// RespondWithSuccess : respond api with success
func RespondWithSuccess(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, message)
}