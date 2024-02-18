package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

var (
	byteCount bool
	lineCount bool
	wordCount bool
	characterCount bool
)

func main() {
	flag.BoolVar(&byteCount, "c", false, "get byte count")
	flag.BoolVar(&lineCount, "l", false, "get line count")
	flag.BoolVar(&wordCount, "w", false, "get word count")
	flag.BoolVar(&characterCount, "m", false, "get character count")
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

	if byteCount {
		fmt.Printf("    %d", len(b))
	}

	reader := bytes.NewReader(b)
	scanner := bufio.NewScanner(reader)
	if lineCount {
		lc := 0
		for scanner.Scan() {
			lc++
		}
		fmt.Printf("    %d", lc)
	}
	
	if wordCount {
		wc := 0
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wc++
		}
		fmt.Printf("    %d", wc)
	}

	if characterCount {
		cc := 0
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			cc++
		}
		fmt.Printf("    %d", cc)
	}

	fmt.Printf(" %s\n", filename)
}