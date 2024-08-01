package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateISBN() string {
	// Seed the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate random parts of the ISBN
	part1 := r.Intn(1000)   // 3-digit part
	part2 := r.Intn(10)     // 1-digit part
	part3 := r.Intn(100000) // 5-digit part
	part4 := r.Intn(1000)   // 3-digit part
	part5 := r.Intn(10)     // 1-digit part

	// Format and return the ISBN
	isbn := fmt.Sprintf("%03d-%d-%05d-%03d-%d",
		part1,
		part2,
		part3,
		part4,
		part5,
	)

	return isbn
}
