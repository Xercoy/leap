package main

import (
	_ "github.com/xercoy/leap"
	"log"
	"os"
)

func main() {
	args := os.Args
	keyword := args[1]

	switch {
	case keyword == "new":
		if len(args) == 4 {
			log.Printf("Adding %s as a place with alias %s.\n",
				args[2], args[3])
		} else {
			log.Printf("Error: must provide exactly one new space and alais, separated by a space.")
		}

	case keyword == "list":
		log.Printf("Showing all current places and their aliases.")

	case keyword == "rm":
		if len(args) == 2 {
			log.Printf("Error: must provide an alias to be deleted.\n")
		} else if len(args) > 4 {
			log.Printf("Error: only one alias can be deleted at a time.\n")
		} else if len(args) == 3 {
			log.Printf("Attempting to delete place with alias %s.\n", args[2])
		}
	}
}
