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
   %s [-v] [-n] [-i package] [-i ...] [-d definition-code] [-b begin-code] [-e end-code] -r codes\n`, os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {
	key := true
	keys := ""
	withLoop := false
	printCode := false
	var ims []string
	var define = ""
	var begin = ""
	var end = ""
	var body = ""
	for i, a := range os.Args {
		if i == 0 {
			continue // skip process name
		}
	AfterOneShot:
		if key {
			if '-' != a[0] {
				log.Fatalf("is not key: %d, %s", i, a)
				usage()
			}
			keys = a
		} else {
			switch keys {
			case "-h":
				usage() //with death
			case "-i":
				ims = append(ims, a)
			case "-d":
				define = a
			case "-b":
				begin = a
			case "-e":
				end = a
			case "-r":
				body = a
			case "-n":
				withLoop = true
				key = !key
				goto AfterOneShot
			case "-v":
				printCode = true
				key = !key
				goto AfterOneShot
			default:
				log.Fatalf("unknown key: %d, %s", i, a)
				usage()
			}
		}
		key = !key
	}

	err := lib.Run(define, begin, body, end, withLoop, printCode, ims...)
	if err != nil {
		log.Fatal(err)
	}
}
