package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	characterCount bool
)



func main() {
	flag.BoolVar(&characterCount, "c", false, "get character count")
	flag.Parse()

	// get filename
	filename := flag.Arg(0)
	if len(filename) == 0 {
		fmt.Println("then do something")
	}

	// ReadFile
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	if characterCount {
		fmt.Printf("    %d", len(b))
	}

	fmt.Printf(" %s\n", filename)
}