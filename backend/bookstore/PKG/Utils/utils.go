package utils

import (
	"time"
)

// FormatDate formats a date into a specific format.
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// IsValidISBN checks if a given string is a valid ISBN.
func IsValidISBN(isbn string) bool {
	// Implement ISBN validation logic
	return true // Replace with actual validation
}

// ErrorHandler handles errors gracefully.
func ErrorHandler(err error) {
	// Implement error handling logic
}

// Constants
const (
	MaxBookTitleLength = 100
	DefaultPageSize    = 10
)
