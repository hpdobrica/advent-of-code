package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("joka-input.txt")
	check(err)
	defer inputFile.Close()

	overlappingPairs := getOverlappingPairCount(inputFile)

	fmt.Println("There are", overlappingPairs, "overlapping pairs")

}

func getOverlappingPairCount(inputFile *os.File) int {
	overlappingPairs := 0

	forLineOfFile(inputFile, func(line string) {
		first, second, third, fourth := extractData(line)
		if isOverlapping(first, second, third, fourth) {
			overlappingPairs++
		}
		fmt.Println(first, second, third, fourth)
	})

	return overlappingPairs
}

func isOverlapping(first int, second int, third int, fourth int) bool {
	return (first <= third && second >= fourth) || (third <= first && fourth >= second)
}

func extractData(line string) (int, int, int, int) {

	re := regexp.MustCompile(`(\d*)-(\d*),(\d*)-(\d*)`)
	match := re.FindStringSubmatch(line)
	first, _ := strconv.Atoi(match[1])
	second, _ := strconv.Atoi(match[2])
	third, _ := strconv.Atoi(match[3])
	fourth, _ := strconv.Atoi(match[4])

	return first, second, third, fourth
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
