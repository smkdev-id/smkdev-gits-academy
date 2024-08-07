package common

import (
	"time"

	"golang.org/x/exp/rand"
)

// randomGeneratedDate generates a random date between the provided start and end dates.
func RandomGeneratedDate(start, end time.Time) time.Time {
	// Calculate the duration between start and end dates
	duration := end.Sub(start)
	// Generate a random number of nanoseconds between 0 and the duration
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	// Add the random duration to the start date
	return start.Add(randomDuration)
}
