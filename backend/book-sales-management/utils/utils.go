package utils

import (
    "log"
)

// LogError logs an error message
func LogError(err error) {
    if err != nil {
        log.Println("ERROR:", err)
    }
}

// LogInfo logs an informational message
func LogInfo(message string) {
    log.Println("INFO:", message)
}
