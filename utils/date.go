package utils

import (
	"time"
)

// GetCurrentDateTimeString returns the current date and time as a formatted string, e.g. "20250101-08:00:12"
func GetCurrentDateTimeString() string {
	return time.Now().Format("20060102150405")
}
