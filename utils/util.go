package utils

import (
	"bufio"
	"fmt"
	"os"
)

func MakeDir(dirName string) {
	err := os.MkdirAll(dirName, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create a dir: %s", err)
	}
}

func WriteFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not create file: %s", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not write to file: %s", err)
	}
	err = writer.Flush()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not flush buffer: %s", err)
	}
}
