package time

import (
	"errors"
	"sort"
	"time"
)

// Series time series methods
type Series struct {
	ts []time.Time
}

// NewSeries return a time series
func NewSeries(ts []time.Time) Series {
	return Series{ts}
}

// IsSameDay returns true if the time series is on the same day
func (s Series) IsSameDay() bool {
	day := s.ts[0].YearDay()
	for _, t := range s.ts {
		if t.YearDay() != day {
			return false
		}
	}
	return true
}

// IsFreeDay return true if day is toll free
func (s Series) IsFreeDay() (bool, error) {
	t := s.ts[0]
	switch t.Year() {
	case 2017:
		return isFree(t, PublicHolidays2017, FreeDays2017), nil
	default:
		return false, errors.New("Wrong year")
	}
}

func isFree(t time.Time, holidays, freeDays []string) bool {
	d := date{Time: t, publicHolidays: holidays, freeDays: freeDays}
	if d.isWeekend() {
		return true
	}
	if d.isPublicHoliday() {
		return true
	}
	if d.isFreeDay() {
		return true
	}

	return false
}

// Partition on duration
func (s Series) Partition(d time.Duration) map[int][]time.Time {
	ts := s.ts
	sort.Sort(ByAsc(ts))

	key := 1
	p := make(map[int][]time.Time)
	start := ts[0]
	for i, t := range ts {
		if t.Sub(start) <= d {
			p[key] = append(p[key], t)
			continue
		}
		start = ts[i]
		key++
		p[key] = append(p[key], t)
	}

	return p
}
