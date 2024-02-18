package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

var (
	characterCount bool
	lineCount bool
)

func main() {
	flag.BoolVar(&characterCount, "c", false, "get character count")
	flag.BoolVar(&lineCount, "l", false, "get line count")
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

	if lineCount {
		reader := bytes.NewReader(b)
		scanner := bufio.NewScanner(reader)
		lc := 0
		for scanner.Scan() {
			lc++
		}
		fmt.Printf("    %d", lc)
	}
	

	fmt.Printf(" %s\n", filename)
}