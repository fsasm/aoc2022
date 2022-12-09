package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota + 1
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

		if false {
			me := parseToken(tokens[1])
			// part 1
			totalScore += me

			if me == opponent {
				// draw
				totalScore += 3
			}

			if ((me == rock) && (opponent == scissor)) ||
				((me == paper) && (opponent == rock)) ||
				((me == scissor) && (opponent == paper)) {
				// win
				totalScore += 6
			}
		} else {
			// part 2
			me := tokens[1]

			if me == "X" {
				// need to lose
				switch opponent {
				case rock:
					totalScore += scissor
				case paper:
					totalScore += rock
				case scissor:
					totalScore += paper
				}
			} else if me == "Y" {
				// draw
				totalScore += opponent
				totalScore += 3
			} else if me == "Z" {
				// need to win
				totalScore += 6
				switch opponent {
				case rock:
					totalScore += paper
				case paper:
					totalScore += scissor
				case scissor:
					totalScore += rock
				}
			}
		}
	}

	fmt.Printf("Total Score: %d\n", totalScore)
}
