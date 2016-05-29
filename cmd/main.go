package main

import (
	"fmt"
	"github.com/xercoy/leap"
	"os"
)

func main() {
	lI := leap.NewLeapInfo(leap.DefaultCfgPath)
	args := os.Args
	var err error

	if len(args) <= 1 {
		fmt.Printf("./")
		return
	}

	switch {
	case (len(args) == 2) && (args[1] == "list"):
		// Write the places to output.
		fmt.Printf("%v", lI.Places)

	case (len(args) == 4) && (args[1] == "add"):
		err = lI.AddPlace(args[2], args[3])
		if err != nil {
			panic(err)
		}

		//fmt.Printf("%v", lI.Places)
	default:
		// Regard args[1] as an alias and attempt to jump.
		fmt.Printf("%v", lI.Leap(args[1]))
	}
}
