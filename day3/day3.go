package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func duplicateLetters(a, b string) (result string) {
	for _, letter := range a {
		if strings.ContainsRune(b, letter) {
			result += string(letter)
		}
	}

	return result
}

func score(r byte) int {
	if ('a' <= r) && (r <= 'z') {
		return int(r-'a') + 1
	}
	if ('A' <= r) && (r <= 'Z') {
		return int(r-'A') + 27
	}

	return 0
}

func processLine(line string) int {
	length := utf8.RuneCountInString(line)
	if (length % 2) != 0 {
		fmt.Println("Lines must have even number of characters")
	}

	a := line[:length/2]
	b := line[length/2:]

	dup := duplicateLetters(a, b)
	return score(dup[0])
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Couldn't open input file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		totalScore += processLine(line)
	}

	fmt.Println("Total score: ", totalScore)
}
