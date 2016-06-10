package main

import (
	"fmt"
	"github.com/xercoy/leap"
	"log"
	"os"
)

func main() {
	lI := leap.NewLeapInfo(leap.DefaultCfgPath)
	args := os.Args
	var err error

	helpText := `Usage:
List all Places          : leap list
Add a Place to leap to   : leap add  <directory> <alias>
Display help text        : leap help
Remove a Place           : leap rm   <alias>`

	switch {
	case ((len(args) == 2) && (args[1] == "help")) || (len(args) == 1):
		fmt.Println(helpText)

	case ((len(args) == 3) && (args[1] == "rm")):

		err = lI.RemovePlace(args[2])
		if err != nil {
			fmt.Printf(err.Error())
			break
		}

		log.Printf("Place with alias %s has successfully been removed.", args[2])

	case (len(args) == 2) && (args[1] == "list"):
		// Write the places to output.
		fmt.Printf("%s", lI.StrOfPlaces())

	case (len(args) == 4) && (args[1] == "add"):
		err = lI.AddPlace(args[2], args[3])
		if err != nil {
			fmt.Printf(err.Error())
			break
		}

		log.Printf("Place with alias %s has successfully been added.", args[3])

	default:
		// Regard args[1] as an alias and attempt to jump.
		fmt.Printf("%v", lI.Leap(args[1]))
	}
}
