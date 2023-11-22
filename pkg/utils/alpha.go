package utils

import (
	"math/rand"
	"time"
)

var ALPHABET []string = []string{
	"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z",
}

func GenerateAlphaMap() map[string]bool {
	// Note: rand.Seed(time.Now().UnixNano()) - DEPRECATED since 1.20
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	m := make(map[string]bool)

	for _, k := range ALPHABET {
		m[k] = rng.Intn(2) == 0
	}

	return m
}
