package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	priority := calculateTotalItemPriority(file)

	fmt.Println("Total item priority is", priority)
}

func calculateTotalItemPriority(file *os.File) int {

	totalPriority := 0

	forLineOfFile(file, func(line string) {
		misplacedItem := findMisplacedItem(line)
		totalPriority += calculateItemPriority(misplacedItem)

	})

	return totalPriority
}

func calculateItemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - ('a' - 1)
	}

	if item >= 'A' && item <= 'Z' {
		capitalAScore := 27
		return int(item) - ('A' - capitalAScore)
	}
	panic("Invalid item, cant calculate priority")

}

func findMisplacedItem(rucksack string) rune {
	firstCompartment := getUniqueItemsInCompartment(rucksack[0 : len(rucksack)/2])
	secondCompartment := getUniqueItemsInCompartment(rucksack[len(rucksack)/2:])

	for _, firstItem := range firstCompartment {
		for _, secondItem := range secondCompartment {
			if firstItem == secondItem {
				return firstItem
			}
		}
	}
	panic("No misplaced item found")
}

func getUniqueItemsInCompartment(compartment string) []rune {
	uniquesMap := map[rune]bool{}

	for _, item := range compartment {
		if uniquesMap[item] != true {
			uniquesMap[item] = true
		}
	}

	uniqueChars := make([]rune, len(uniquesMap))

	i := 0
	for key, _ := range uniquesMap {
		uniqueChars[i] = key
		i++
	}

	return uniqueChars
}

// general utils

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

// FnVRRsVdSnSnFSRqTVdq BBDBhrDdmcddMcMQMhzm
