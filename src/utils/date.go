package utils

import "time"

// parseDate parses a date string in the format "dd/mm/yyyy" and returns time.Time
// If the date string is invalid, it returns a empty time.Time
func ParseDate(date string) (time.Time) {
	parsedDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		return time.Time{}
	}
	return parsedDate
}

// Returns the current time in UTC-3
func TimeNow() time.Time {
	return time.Now().UTC().Add(-3 * time.Hour)
}