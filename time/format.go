package time

import (
	"strings"
	"time"
)

// DateToTime returns time.Time
//
//  Usage:
//         t, err := DateToTime("2017-01-01")
//
func DateToTime(d string) (time.Time, error) {
	return time.Parse(time.RFC3339, d+"T00:00:00Z")
}

// ToSeries returns []time.Time for the given day
//
//  Usage:
//        t := time.Now()
//        ts, err := toSeries("05:56,06:10", t)
//
func ToSeries(ts string, day time.Time) ([]time.Time, error) {
	tsplit := strings.Split(ts, ",")
	ta := make([]time.Time, len(tsplit))

	date := day.Format(time.RFC3339)
	for i, t := range tsplit {
		t, err := ToTime(date[:10], t+":00")
		if err != nil {
			return nil, err
		}
		ta[i] = t
	}
	return ta, nil
}

// ToTime returns time.Time
//
// Usage:
//
//        t, err := ToTime("2017-01-01", "10:10:59")
//
func ToTime(date, t string) (time.Time, error) {
	return time.Parse(time.RFC3339, date+"T"+t+"Z")
}
