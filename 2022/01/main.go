package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	first, topThree := elfCaloriesCalcuator(file)
	fmt.Println("Elf with most calories has", first, "calories", "while the top three together have", topThree, "calories")

}

// returns how much calories the first elf has, and how much calories the top three have together
func elfCaloriesCalcuator(file *os.File) (int, int) {
	currentCalories := 0

	first := 0 // largest
	second := 0
	third := 0 // smallest

	forLineOfFile(file, func(line string) {
		if line == "" {
			if currentCalories > third {
				includeInTopThree(currentCalories, &first, &second, &third)
			}
			currentCalories = 0
			return
		}

		calories, err := strconv.Atoi(line)
		check(err)
		currentCalories += calories

	})

	return first, first + second + third

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func forLineOfFile(file *os.File, fn func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
}

func includeInTopThree(current int, first, second, third *int) {
	if current > *first {
		*third = *second
		*second = *first
		*first = current
		return
	}
	if current > *second {
		*third = *second
		*second = current
		return
	}
	if current > *third {
		*third = current
		return
	}

}
