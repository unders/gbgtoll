package prog

import (
	"github.com/unders/gbgtoll/time"
	"github.com/unders/gbgtoll/toll"
	"github.com/unders/gbgtoll/vehicle"
)

// CalcTollFee returns the toll fee
func CalcTollFee(v vehicle.Type, ts time.Series,
	maxAmount int) Result {

	if toll.VehicleIsFree(v) {
		return Result{Fee: 0, Type: TollFreeVehicle}
	}

	if !ts.IsSameDay() {
		return Result{Fee: 0, Type: NotSameDay}
	}

	free, err := ts.IsFreeDay()
	if err != nil {
		return Result{Fee: 0, Type: WrongYear}
	}

	if free {
		return Result{Fee: 0, Type: TollFreeDay}
	}

	fee := toll.CalcFee(ts, maxAmount)
	return Result{Fee: fee, Type: Fee}
}
