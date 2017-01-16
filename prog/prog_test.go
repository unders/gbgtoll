package prog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/unders/gbgtoll/prog"
	"github.com/unders/gbgtoll/test"
	ttime "github.com/unders/gbgtoll/time"
	"github.com/unders/gbgtoll/vehicle"
)

func TestProg(t *testing.T) {
	t.Run("No toll for free vehicles", noTollForFreeVehicles)
	t.Run("Time series must be on the same day", timeSeriesMustBeOnSameDay)
	t.Run("Year 2017 has an implementation", yearWorks("2017-01-17"))
	t.Run("No toll on weekends", noTollOnWeekends)
	f := noTollOnTheseDays("2017", ttime.PublicHolidays2017)
	t.Run("Year 2017 has no toll on public holidays", f)
	f = noTollOnTheseDays("2017", ttime.FreeDays2017)
	t.Run("Year 2017 has no toll on free days", f)
	t.Run("Toll fee on a given time", calcTollFee)
	t.Run("Toll fee has a max fee per day", calcTollMax(60))
}

func noTollForFreeVehicles(t *testing.T) {
	want := prog.Result{Type: prog.TollFreeVehicle, Fee: 0}
	vehicles := []vehicle.Type{vehicle.Emergency, vehicle.Bus, vehicle.Diplomat}
	t.Logf("toll free vehicles: %#v", vehicles)

	for _, v := range vehicles {
		got := prog.CalcTollFee(v, ttime.Series{}, 60)
		if want != got {
			t.Errorf("\nWant: %#v\n Got: %#v", want, got)
		}
	}
}
func timeSeriesMustBeOnSameDay(t *testing.T) {
	ts := []time.Time{
		test.OnDay(t, "2017-01-02")("06:40"),
		test.OnDay(t, "2017-01-03")("06:40"),
	}
	want := prog.Result{Type: prog.NotSameDay, Fee: 0}
	got := prog.CalcTollFee(vehicle.Car, ttime.NewSeries(ts), 60)
	if want != got {
		t.Errorf("\nWant: %#v\n Got: %#v", want, got)
	}
}

func noTollOnWeekends(t *testing.T) {
	want := prog.Result{Type: prog.TollFreeDay, Fee: 0}

	ts := []time.Time{test.Saturday(t)}
	got := prog.CalcTollFee(vehicle.Pickup, ttime.NewSeries(ts), 60)
	if want != got {
		t.Errorf("\nWant: %#v\n Got: %#v", want, got)
	}

	ts = []time.Time{test.Sunday(t)}
	got = prog.CalcTollFee(vehicle.Pickup, ttime.NewSeries(ts), 60)
	if want != got {
		t.Errorf("\nWant: %#v\n Got: %#v", want, got)
	}
}
func yearWorks(day string) func(*testing.T) {
	want := prog.Result{Type: prog.Fee, Fee: 9}

	return func(t *testing.T) {
		ts := []time.Time{test.OnDay(t, day)("06:29")}

		got := prog.CalcTollFee(vehicle.Pickup, ttime.NewSeries(ts), 60)
		if want != got {
			t.Errorf("\nWant: %#v\n Got: %#v", want, got)
		}
	}
}

func noTollOnTheseDays(year string, days []string) func(t *testing.T) {
	want := prog.Result{Type: prog.TollFreeDay, Fee: 0}

	return func(t *testing.T) {
		times := test.ToTimes(t, year, days, "06:29")

		t.Logf("Free days: %v", days)

		for _, tt := range times {
			ts := ttime.NewSeries([]time.Time{tt})
			got := prog.CalcTollFee(vehicle.Pickup, ts, 60)
			if want != got {
				t.Errorf("\nWant: %#v\n Got: %#v", want, got)
			}
		}
	}
}

// Result returns prog.Result
func Result(fee int) prog.Result {
	return prog.Result{Type: prog.Fee, Fee: fee}
}

func calcTollFee(tt *testing.T) {
	t := test.OnDay(tt, "2017-01-03")

	tests := []struct {
		want prog.Result
		t    time.Time
	}{
		{want: Result(0), t: t("05:59")},
		{want: Result(9), t: t("06:00")},
		{want: Result(9), t: t("06:29")},
		{want: Result(16), t: t("06:30")},
		{want: Result(16), t: t("06:59")},
		{want: Result(22), t: t("07:00")},
		{want: Result(22), t: t("07:59")},
		{want: Result(16), t: t("08:00")},
		{want: Result(16), t: t("08:29")},
		{want: Result(9), t: t("08:30")},
		{want: Result(9), t: t("14:59")},
		{want: Result(16), t: t("15:00")},
		{want: Result(16), t: t("15:29")},
		{want: Result(22), t: t("15:30")},
		{want: Result(22), t: t("16:59")},
		{want: Result(16), t: t("17:00")},
		{want: Result(16), t: t("17:59")},
		{want: Result(9), t: t("18:00")},
		{want: Result(9), t: t("18:29")},
		{want: Result(0), t: t("18:30")},
	}

	for _, tc := range tests {
		fee := fmt.Sprintf(" %s", tc.want)
		tt.Run(tc.t.String()+fee, func(tt *testing.T) {
			ts := ttime.NewSeries([]time.Time{tc.t})

			got := prog.CalcTollFee(vehicle.Pickup, ts, 60)
			if tc.want != got {
				tt.Errorf("\nWant: %#v\n Got: %#v\n", tc.want, got)
			}
		})
	}
}

func calcTollMax(max int) func(*testing.T) {
	return func(t *testing.T) {
		tt := test.OnDay(t, "2017-01-03")
		ts69 := ttime.NewSeries([]time.Time{
			tt("07:00"), // 22 kr
			tt("08:29"), // 16 kr
			tt("15:30"), // 22 kr
			tt("18:00"), // 9 kr
		})
		want := prog.Result{Type: prog.Fee, Fee: max}

		t.Run("Max"+want.String(), func(tt *testing.T) {
			got := prog.CalcTollFee(vehicle.Pickup, ts69, max)
			if want != got {
				t.Errorf("\nWant: %#v\n Got: %#v\n", want, got)
			}
		})
	}
}
