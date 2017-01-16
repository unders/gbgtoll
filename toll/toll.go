package toll

import (
	"math"
	"time"

	ttime "github.com/unders/gbgtoll/time"
	"github.com/unders/gbgtoll/vehicle"
)

// VehicleIsFree returns true if Vehicle is toll free
func VehicleIsFree(v vehicle.Type) bool {
	switch v {
	case vehicle.Car, vehicle.Pickup:
		return false
	default:
		return true
	}
}

// CalcFee return the toll fee
func CalcFee(s ttime.Series, max int) int {
	tp := s.Partition(1 * time.Hour)

	fee := 0
	for _, ts := range tp {
		fee += maxFee(ts)
	}

	return int(math.Min(float64(max), float64(fee)))
}

func maxFee(ts []time.Time) int {
	fee := 0
	for _, t := range ts {
		p := calcFee(t)
		if p > fee {
			fee = p
		}
	}
	return fee
}

// calcFee returns the toll fee
func calcFee(t time.Time) int {
	switch h, m, _ := t.Clock(); true {
	case h == 6 && m < 30:
		return 9
	case h == 6 && m <= 59:
		return 16
	case h == 7 && m <= 59:
		return 22
	case h == 8 && m < 30:
		return 16
	case h == 8 && m <= 59:
		return 9
	case h > 8 && h < 15:
		return 9
	case h == 15 && m < 30:
		return 16
	case h == 15 && m <= 59:
		return 22
	case h == 16:
		return 22
	case h == 17:
		return 16
	case h == 18 && m < 30:
		return 9
	default:
		return 0
	}
}
