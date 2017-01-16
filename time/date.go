package time

import (
	"strconv"
	"time"
)

type date struct {
	time.Time
	publicHolidays []string
	freeDays       []string
}

// dayMonth returns day-month as a string
//
// Example:
//
//          7-January
//
func (d date) dayMonth() string {
	return strconv.Itoa(d.Time.Day()) + "-" + d.Time.Month().String()
}

// isWeekend returns true if weekend
func (d date) isWeekend() bool {
	day := d.Weekday()
	// Sunday = 0; Saturday = 6
	return day == 0 || day == 6
}

// isPublicHoliday returns true if public holiday
func (d date) isPublicHoliday() bool {
	dayMonth := d.dayMonth()

	for _, day := range d.publicHolidays {
		if dayMonth == day {
			return true
		}
	}

	return false
}

// isFreeDays returns true if free day
func (d date) isFreeDay() bool {
	dayMonth := d.dayMonth()

	for _, day := range d.freeDays {
		if dayMonth == day {
			return true
		}
	}

	return false
}
