package vehicle

import (
	"errors"
	"fmt"
)

// Type type definition
type Type string

// List of vehicle types
const (
	Car       Type = "car"
	Pickup         = "pickup"
	Emergency      = "emergency"
	Bus            = "bus"
	Diplomat       = "diplomat"
)

// Types of different types
var Types = map[Type]Type{
	Car:       Car,
	Pickup:    Pickup,
	Emergency: Emergency,
	Bus:       Bus,
	Diplomat:  Diplomat,
}

// Get returns vehicle.Type
func Get(s string) (Type, error) {
	if v, ok := Types[Type(s)]; ok {
		return v, nil
	}
	msg := fmt.Sprintf("%s is not a vehicle type\n", s)
	return Car, errors.New(msg)
}
