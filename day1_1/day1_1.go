package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	id       int
	calories []int
}

func NewElf(id int) Elf {
	return Elf{id: id, calories: make([]int, 0)}
}

func (elf *Elf) AddCallories(calories int) {
	elf.calories = append(elf.calories, calories)
}

func (elf Elf) TotalCallories() (sum int) {
	for _, callory := range elf.calories {
		sum += callory
	}

	return
}

func main() {
	const FILENAME = "input"

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Println("Couldn't open input file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentId := 0
	currentElf := NewElf(currentId)

	elfs := make([]Elf, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elfs = append(elfs, currentElf)
			currentId++
			currentElf = NewElf(currentId)
		} else {
			callory, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Error while parsing line\n")
				return
			}
			currentElf.AddCallories(callory)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	maximumCalories := 0
	for _, elf := range elfs {
		if maximumCalories < elf.TotalCallories() {
			maximumCalories = elf.TotalCallories()
		}
	}

	fmt.Printf("Highes calories: %d\n", maximumCalories)
}
