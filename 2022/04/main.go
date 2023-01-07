package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	check(err)
	defer inputFile.Close()

	totallyOverlappingPairs, partiallyOverlappingPairs := getOverlappingPairCount(inputFile)

	fmt.Println("There are", totallyOverlappingPairs, "totally overlapping pairs, and", partiallyOverlappingPairs, "partially overlapping pairs.")

}

func getOverlappingPairCount(inputFile *os.File) (int, int) {
	totallyOverlappingPairs := 0
	partiallyOverlappingPairs := 0

	forLineOfFile(inputFile, func(line string) {
		first, second, third, fourth := extractData(line)
		if isTotallyOverlapping(first, second, third, fourth) {
			totallyOverlappingPairs++
			partiallyOverlappingPairs++
		} else if isPartiallyOverlapping(first, second, third, fourth) {
			partiallyOverlappingPairs++
		}
		fmt.Println(first, second, third, fourth)
	})

	return totallyOverlappingPairs, partiallyOverlappingPairs
}

func isTotallyOverlapping(first int, second int, third int, fourth int) bool {
	return (first <= third && second >= fourth) || (third <= first && fourth >= second)
}

func isPartiallyOverlapping(first int, second int, third int, fourth int) bool {
	return second >= third && first <= fourth
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
