package main

import (
	"fmt"
	"os"

	"github.com/unders/gbgtoll/cli"
	"github.com/unders/gbgtoll/prog"
)

// MaxAmount the max amount to pay per day
const MaxAmount int = 60

func main() {
	arg, usage, err := cli.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		fmt.Println(usage)
		os.Exit(0)
	}

	r := prog.CalcTollFee(arg.Vehicle, arg.Series, MaxAmount)
	fmt.Println(r)
}
