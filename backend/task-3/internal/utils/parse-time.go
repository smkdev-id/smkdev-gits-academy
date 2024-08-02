package utils

import (
	"fmt"
	"time"
)

// ParseTime converts a string in the format dd-MM-yyyy to time.Time
func ParseTime(dateStr string) (time.Time, error) {
	layout := "02-01-2006"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	return t, nil
}

// FormatTime converts a time.Time to string in the format dd-MM-yyyy
func FormatTime(t time.Time) string {
	layout := "02-01-2006"
	return t.Format(layout)
}
