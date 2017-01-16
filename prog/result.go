package prog

import "fmt"

// Type defines different result types
type Type int

// List of result types
const (
	TollFreeVehicle Type = iota
	NotSameDay
	WrongYear
	TollFreeDay
	Fee
)

// Result contains the result from the computation
type Result struct {
	Type Type
	Fee  int
}

func (r Result) String() string {
	switch r.Type {
	case TollFreeVehicle:
		return fmt.Sprint("vehicle is toll free")
	case NotSameDay:
		return fmt.Sprint("time series must be on same day")
	case WrongYear:
		return fmt.Sprint("Year is not implemented")
	case TollFreeDay:
		return fmt.Sprint("Day is toll free")
	case Fee:
		return fmt.Sprintf("Fee %d kr", r.Fee)
	default:
		return "No implemented"
	}
}
