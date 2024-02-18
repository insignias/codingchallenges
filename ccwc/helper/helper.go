package helper

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func ReadFile(filename string) ([]byte) {
	if len(filename) == 0 {
		fmt.Println("then do something")
	}

	// ReadFile
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

func GetByteCount(b [] byte) int {
	return len(b)
}

func getNewScanner(b []byte) *bufio.Scanner {
	reader := bytes.NewReader(b)
	scanner := bufio.NewScanner(reader)

	return scanner
}

func GetLineCount(b []byte) int {
	scanner := getNewScanner(b)
	lc := 0
	for scanner.Scan() {
		lc++
	}
	return lc
}

func GeWordCount(b []byte) int {
	scanner := getNewScanner(b)
	wc := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wc++
	}
	return wc
}

func GetCharacterCount(b []byte) int {
	scanner := getNewScanner(b)
	cc := 0
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		cc++
	}
	return cc
}