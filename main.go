package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kuwa72/gowk/lib"
)

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [-i package] [-i package] [-b begin-code] [-e end-code] -r codes\n`, os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {
	key := true
	keys := ""
	var ims []string
	var begin = ""
	var end = ""
	var body = ""
	for i, a := range os.Args {
		if i == 0 {
			continue // skip process name
		}
		if key {
			if '-' != a[0] {
				log.Fatalf("is not key: %d, %s", i, a)
				usage()
			}
			keys = a
		} else {
			switch keys {
			case "-i":
				ims = append(ims, a)
			case "-b":
				begin = a
			case "-e":
				end = a
			case "-r":
				body = a
			default:
				log.Fatalf("unknown key: %d, %s", i, a)
				usage()
			}
		}
		key = !key
	}

	err := lib.Run(begin, body, end, ims...)
	if err != nil {
		log.Fatal(err)
	}
}
