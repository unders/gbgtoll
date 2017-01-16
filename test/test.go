package test

import (
	"strings"
	"testing"
	"time"

	ttime "github.com/unders/gbgtoll/time"
)

// OnDay return a function that takes a string and
// returns a time.Time
func OnDay(t *testing.T, day string) func(string) time.Time {
	d := day
	return func(tStr string) time.Time {
		tt, err := ttime.ToTime(d, tStr+":00")
		if err != nil {
			t.Fatalf("\nWant nil\n Got: %s", err)
		}
		return tt
	}
}

// Sunday returns time.Time that is a Sunday
func Sunday(t *testing.T) time.Time {
	return OnDay(t, "2017-01-22")("06:40")
}

// Saturday returns time.Time that is a Saturday
func Saturday(t *testing.T) time.Time {
	return OnDay(t, "2017-01-21")("06:40")
}

// ToTimes returns time.Time
func ToTimes(t *testing.T, year string, dayMonth []string, tt string) []time.Time {
	ts := make([]time.Time, len(dayMonth))

	for i, dm := range dayMonth {
		sp := strings.Split(dm, "-")
		day, month := day(sp[0]), month[sp[1]]
		ts[i] = OnDay(t, year+"-"+month+"-"+day)(tt)
	}

	return ts
}

var month = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

func day(day string) string {
	if len(day) == 1 {
		return "0" + day
	}
	return day
}
