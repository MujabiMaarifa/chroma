package cmd

import (
	"fmt"
	"os"
)

func generate(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%s", file)
}
