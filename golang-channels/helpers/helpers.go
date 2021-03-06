package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
