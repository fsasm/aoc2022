package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	id       int
	calories []int
}

func NewElf(id int) Elf {
	return Elf{id: id, calories: make([]int, 0)}
}

func (elf *Elf) AddCalorie(calorie int) {
	elf.calories = append(elf.calories, calorie)
}

func (elf Elf) TotalCalories() (sum int) {
	for _, calorie := range elf.calories {
		sum += calorie
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
	calories := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elfs = append(elfs, currentElf)
			calories = append(calories, currentElf.TotalCalories())
			currentId++
			currentElf = NewElf(currentId)
		} else {
			calory, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Error while parsing line\n")
				return
			}
			currentElf.AddCalorie(calory)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Println("Top calories: ", calories[:5])
	fmt.Println("Sum of top 3: ", (calories[0] + calories[1] + calories[2]))
}
