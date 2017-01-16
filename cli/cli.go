package cli

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"time"

	ttime "github.com/unders/gbgtoll/time"
	"github.com/unders/gbgtoll/vehicle"
)

const usage = `
Usage:

    gbgtoll vehicle -d=date -t='time series for that day'


Vehicles:
    car            #  private car
    pickup
    emergency      #  emergency vehicle
    bus
    diplomat       #  diplomat registered cars

Flags:
    -d             #  the date of the time series;
                      format: 2017-01-01
    -t             #  time series for when the vehicle passed tolls;
                      format: 05:56,06:10
    -help          #  show this help message


Examples:
    gbgtoll -help
    gbgtoll car -d=2017-01-01 -t=05:56,06:10
`

// Arg data parsed from the command line
type Arg struct {
	Vehicle vehicle.Type
	Series  []time.Time
}

// Parse returns the argument from the command line
//
// Usage:
//        arg, usage, err := cli.Parse(os.Args)
//
func Parse(args []string) (Arg, string, error) {
	arg := Arg{}
	if len(args) < 2 {
		return arg, usage, errors.New("")
	}

	progName := args[0]
	f := flag.NewFlagSet(progName, flag.ContinueOnError)
	buf := &bytes.Buffer{}
	f.SetOutput(buf)

	var err error
	if arg.Vehicle, err = vehicle.Get(args[1]); err != nil {
		return arg, usage, err
	}

	var (
		t string
		d string
	)
	f.StringVar(&t, "t", "", "")
	f.StringVar(&d, "d", "", "")

	if err := f.Parse(args[2:]); err != nil {
		return arg, usage, err
	}

	date, err := ttime.DateToTime(d)
	if err != nil {
		msg := fmt.Sprint("\ninvalid flag argument: -d=", d)
		return arg, usage, errors.New(msg)
	}

	if arg.Series, err = ttime.ToSeries(t, date); err != nil {
		msg := fmt.Sprint("\ninvalid flag argument: -t=", t)
		return arg, usage, errors.New(msg)
	}

	return arg, usage, nil
}
