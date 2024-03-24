package helper

import (
	"fmt"
	"io"
	"os"
)


func ReadFromFileOrStdin(filename string) ([]byte) {
	// ReadFile
	b, err := os.ReadFile(filename)
	if err != nil {
		b, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
		}
	}

	return b
}