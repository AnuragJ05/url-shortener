package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// Create a new random number generator with the current time as the seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]rune, length)
	for i := range result {
		result[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(result)
}
