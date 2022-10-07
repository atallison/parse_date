package parse_date

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {

	t.Run("Input 'today' to return today's date", func(t *testing.T) {
		output := parse("today")
		actual_year, actual_month, actual_day := output.Date()
		expected_year, expected_month, expected_day := time.Now().Date()
		if actual_year != expected_year || actual_month != expected_month || actual_day != expected_day {
			t.Log("output should be today's date")
			t.Fail()
		}
	})

	t.Run("Input 'tomorrow' to return tomorrow's date", func(t *testing.T) {
		output := parse("tomorrow")
		actual_year, actual_month, actual_day := output.Date()
		expected_year, expected_month, expected_day := time.Now().Add(time.Hour * 24 * 1).Date()
		if actual_year != expected_year || actual_month != expected_month || actual_day != expected_day {
			t.Log("output should be today's date")
			t.Fail()
		}
	})

}
