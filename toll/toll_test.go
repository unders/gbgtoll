package toll

import (
	"testing"
	"time"

	"github.com/unders/gbgtoll/test"
	ttime "github.com/unders/gbgtoll/time"
)

func TestCalcFee(tt *testing.T) {
	t := test.OnDay(tt, "2017-01-03")
	ts0 := []time.Time{t("04:59"), t("05:59"), t("18:30")}
	ts9 := []time.Time{t("06:01"), t("06:29")}  // same hour
	ts9b := []time.Time{t("09:30"), t("10:30")} // same hour
	ts18 := []time.Time{t("06:01"), t("18:01")}
	ts18b := []time.Time{t("09:30"), t("10:31")}
	//  same hour + 1
	ts18c := []time.Time{t("09:30"), t("10:30"), t("10:31")}
	ts16 := []time.Time{t("06:30"), t("06:35")} // same hour
	ts32 := []time.Time{t("06:30"), t("08:29")}
	// Same hour
	ts22 := []time.Time{t("07:00"), t("07:15"), t("07:30"), t("07:59")}
	// same hour takes highest fee
	ts22b := []time.Time{t("06:58"), t("07:15"), t("07:30"), t("07:58")}

	// checking that sorting of times works
	ts44 := []time.Time{t("07:59"), t("06:58"), t("07:15"), t("07:30")}

	tests := []struct {
		want int
		max  int
		ts   ttime.Series
	}{
		{want: 0, ts: ttime.NewSeries(ts0), max: 60},
		{want: 9, ts: ttime.NewSeries(ts9), max: 60},
		{want: 9, ts: ttime.NewSeries(ts9b), max: 60},
		{want: 18, ts: ttime.NewSeries(ts18), max: 60},
		{want: 18, ts: ttime.NewSeries(ts18b), max: 60},
		{want: 18, ts: ttime.NewSeries(ts18c), max: 60},
		{want: 16, ts: ttime.NewSeries(ts16), max: 60},
		{want: 32, ts: ttime.NewSeries(ts32), max: 60},
		{want: 22, ts: ttime.NewSeries(ts22), max: 60},
		{want: 22, ts: ttime.NewSeries(ts22b), max: 60},
		{want: 44, ts: ttime.NewSeries(ts44), max: 60},
		{want: 42, ts: ttime.NewSeries(ts44), max: 42},
	}

	for _, tc := range tests {
		tt.Run("", func(tt *testing.T) {
			got := CalcFee(tc.ts, tc.max)
			if tc.want != got {
				tt.Errorf("\nWant: %d\n Got: %d\n", tc.want, got)
			}
		})
	}
}

func TestMaxFee(tt *testing.T) {
	t := test.OnDay(tt, "2017-01-03")
	ts0 := []time.Time{t("04:59"), t("05:59"), t("18:30")}
	ts9 := []time.Time{t("11:35"), t("06:29")}
	ts16 := []time.Time{t("05:59"), t("06:30")}
	ts22 := []time.Time{t("05:59"), t("06:30"), t("07:59")}
	ts22b := []time.Time{t("06:30"), t("07:59"), t("18:00")}

	tests := []struct {
		want int
		ts   []time.Time
	}{
		{want: 0, ts: ts0},
		{want: 9, ts: ts9},
		{want: 16, ts: ts16},
		{want: 22, ts: ts22},
		{want: 22, ts: ts22b},
	}

	for _, tc := range tests {
		tt.Run("", func(tt *testing.T) {
			got := maxFee(tc.ts)
			if tc.want != got {
				tt.Errorf("\nWant: %d\n Got: %d\n", tc.want, got)
			}
		})
	}
}

func TestCalcFeeB(tt *testing.T) {
	t := test.OnDay(tt, "2017-01-03")

	tests := []struct {
		want int
		t    time.Time
	}{
		{want: 0, t: t("05:59")},
		{want: 9, t: t("06:00")},
		{want: 9, t: t("06:29")},
		{want: 16, t: t("06:30")},
		{want: 16, t: t("06:59")},
		{want: 22, t: t("07:00")},
		{want: 22, t: t("07:59")},
		{want: 16, t: t("08:00")},
		{want: 16, t: t("08:29")},
		{want: 9, t: t("08:30")},
		{want: 9, t: t("09:30")},
		{want: 9, t: t("11:35")},
		{want: 9, t: t("14:59")},
		{want: 16, t: t("15:00")},
		{want: 16, t: t("15:29")},
		{want: 22, t: t("15:30")},
		{want: 22, t: t("16:59")},
		{want: 16, t: t("17:00")},
		{want: 16, t: t("17:59")},
		{want: 9, t: t("18:00")},
		{want: 9, t: t("18:29")},
		{want: 0, t: t("18:30")},
	}

	for _, tc := range tests {
		tt.Run(tc.t.String(), func(tt *testing.T) {
			got := calcFee(tc.t)
			if tc.want != got {
				tt.Errorf("\nWant: %d\n Got: %d\n", tc.want, got)
			}
		})
	}
}
