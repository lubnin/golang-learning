package main

import (
	"fmt"
	"io"
	"os"
)

// to run: go run main.go file.txt
func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
