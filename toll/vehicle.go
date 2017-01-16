package toll

import (
	"errors"
	"fmt"
)

// Vehicle type definition
type Vehicle string

// List of vehicle types
const (
	Car       Vehicle = "car"
	Pickup            = "pickup"
	Emergency         = "emergency"
	Bus               = "bus"
	Diplomat          = "diplomat"
)

// Vehicles of different types
var Vehicles = map[Vehicle]Vehicle{
	Car:       Car,
	Pickup:    Pickup,
	Emergency: Emergency,
	Bus:       Bus,
	Diplomat:  Diplomat,
}

// LookupVehicle returns Vehicle
func LookupVehicle(s string) (Vehicle, error) {
	if v, ok := Vehicles[Vehicle(s)]; ok {
		return v, nil
	}
	msg := fmt.Sprintf("%s is not a vehicle type\n", s)
	return Car, errors.New(msg)
}
