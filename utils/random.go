package utils

import (
	"math/rand"
)

func GetRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
