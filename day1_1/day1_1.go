package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const FILENAME = "input"

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("Couldn't open input file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Line: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
