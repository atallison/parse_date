package parse_date

import "time"

func parse(input string) time.Time {
	if input == "today" {
		return time.Now()
	}
	if input == "tomorrow" {
		return time.Now().Add(time.Hour * 24)
	}
	return time.Time{}
}
