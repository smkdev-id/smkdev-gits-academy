package utils

import (
	"fmt"
)

// FormatRP formats an integer as a currency string (Rupiah)
func FormatRP(amount uint) string {
	return fmt.Sprintf("Rp %d", amount)
}

// ErrorResponse is a helper function to generate error responses
func ErrorResponse(message string) map[string]string {
	return map[string]string{"error": message}
}
