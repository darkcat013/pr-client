package utils

import (
	"math/rand"
	"time"

	"github.com/darkcat013/pr-client/config"
)

func SleepFor(t float64) {
	time.Sleep(time.Duration(t) * config.TIME_UNIT)
}

func SleepBetween(min, max int) int {
	randomTime := GetRandomInt(min, max)
	time.Sleep(time.Duration(randomTime) * config.TIME_UNIT)
	return randomTime
}

func SleepOneOf(params ...int) {
	time.Sleep(time.Duration(params[rand.Intn(len(params))]) * config.TIME_UNIT)
}

func GetCurrentTimeFloat() float64 {
	if config.TIME_UNIT >= time.Millisecond && config.TIME_UNIT < time.Second {
		return float64(time.Now().UnixMilli())
	} else {
		return float64(time.Now().Unix())
	}
}
