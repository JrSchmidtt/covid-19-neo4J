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