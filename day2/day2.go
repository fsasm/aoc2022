package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota
	paper
	scissor
)

func parseToken(token string) int {
	switch token {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissor
	default:
		panic("Invalid token")
	}
}

func main() {
	file, err := os.Open("input")

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

		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			fmt.Println("Error while parsing line")
			return
		}

		opponent := parseToken(tokens[0])
		me := parseToken(tokens[1])

		totalScore += me + 1

		if me == opponent {
			// draw
			totalScore += 3
		}

		if ((me == rock) && (opponent == scissor)) ||
			((me == paper) && (opponent == rock)) ||
			((me == scissor) && (opponent == paper)) {
			totalScore += 6
		}
	}

	fmt.Printf("Total Score: %d\n", totalScore)
}
