package utils

import (
	"log"
	"math/rand"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

const (
	SMALL  = "small"
	MEDIUM = "medium"
	LARGE  = "large"
)

var FRUITS = []string{
	"Orange",
	"Apple",
	"Carrot",
	"Watermelon",
}

func GetRandomFruit() (string, int32) {
	size := 1 + rand.Int31n(100)
	idx := rand.Int31n(4)
	return FRUITS[idx], size
}
