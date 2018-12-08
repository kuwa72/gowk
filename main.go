package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kuwa72/gowk/lib"
)

func main() {
	tv, err := lib.Eval(`(100 + 1) * len("hoge")`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tv.Value)

	fmt.Printf("Args: %#v", os.Args)
	tv, err = lib.Eval(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tv.Value)
}
