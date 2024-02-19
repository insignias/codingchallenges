package main

import (
	"flag"
	"fmt"
	"github.com/insignias/codingchallenges/ccwc/helper"
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

	//ReadFile
	b := helper.ReadFromFileOrStdin(filename)

	//byteCount
	bc := helper.GetByteCount(b)

	//lineCount
	lc := helper.GetLineCount(b)

	//wordCount
	wc := helper.GetWordCount(b)

	if byteCount {
		fmt.Printf("    %d", bc)
	}

	if lineCount {
		
		fmt.Printf("    %d", lc)
	}
	
	if wordCount {
		fmt.Printf("    %d", wc)
	}

	if characterCount {
		cc := helper.GetCharacterCount(b)
		fmt.Printf("    %d", cc)
	}

	if flag.NFlag() == 0 {
		fmt.Printf("    %d    %d    %d", lc, wc, bc)
	}

	fmt.Printf(" %s\n", filename)
}