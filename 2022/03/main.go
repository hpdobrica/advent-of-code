package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)

	totalPriority, badgePriority := calculateTotalItemPriorities(file)

	fmt.Println("Total item priority is", totalPriority, ", while the total priority of the group badges is", badgePriority)
}

func calculateTotalItemPriorities(file *os.File) (int, int) {

	totalPriority := 0
	badgePriority := 0

	groupsOfThree := make([]string, 3)
	groupCounter := 0

	forLineOfFile(file, func(line string) {
		misplacedItem := findMisplacedItem(line)
		totalPriority += calculateItemPriority(misplacedItem)

		groupsOfThree[groupCounter] = line
		if groupCounter == 2 {
			groupCounter = 0
			badgePriority += calculateItemPriority(findGroupBadge(groupsOfThree[0], groupsOfThree[1], groupsOfThree[2]))
		} else {
			groupCounter++
		}
	})

	return totalPriority, badgePriority
}

func findGroupBadge(rucksackOne string, rucksackTwo string, rucksackThree string) rune {
	rucksacksUniqueItems := [3][]rune{getUniqueItems(rucksackOne), getUniqueItems(rucksackTwo), getUniqueItems(rucksackThree)}

	itemsMap := map[rune]int{}

	for _, rucksack := range rucksacksUniqueItems {
		for _, item := range rucksack {
			itemsMap[item]++
			if itemsMap[item] == 3 {
				return item
			}
		}
	}

	panic("No group badge found")
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
	firstCompartment := getUniqueItems(rucksack[0 : len(rucksack)/2])
	secondCompartment := getUniqueItems(rucksack[len(rucksack)/2:])

	for _, firstItem := range firstCompartment {
		for _, secondItem := range secondCompartment {
			if firstItem == secondItem {
				return firstItem
			}
		}
	}
	panic("No misplaced item found")
}

func getUniqueItems(collection string) []rune {
	uniquesMap := map[rune]bool{}

	uniques := []rune{}

	for _, item := range collection {
		if uniquesMap[item] != true {
			uniquesMap[item] = true
			uniques = append(uniques, item)
		}
	}

	return uniques
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
