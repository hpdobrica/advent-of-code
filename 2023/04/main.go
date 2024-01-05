package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {

	start := time.Now()

	file, err := os.Open("./input.txt")
	util.PanicIfErr(err)

	sumOfPoints := processInput(file)

	fmt.Println("Total point worth of cards is ", sumOfPoints)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)

}

func processInput(file *os.File) int {
	sumOfPoints := 0
	util.ForLineOfFile(file, func(line string) {
		sumOfCardPoints := 0
		winningNumbers, elfsNumbers := parseLine(line)
		fmt.Println(winningNumbers, elfsNumbers)

		winningMap := make(map[int]int)

		for _, winningNum := range winningNumbers {
			_, exists := winningMap[winningNum]
			if !exists {
				winningMap[winningNum] = 0
			}
		}

		for _, elfsNum := range elfsNumbers {
			_, exists := winningMap[elfsNum]
			if exists {
				winningMap[elfsNum] = winningMap[elfsNum] + 1

				if sumOfCardPoints == 0 {
					sumOfCardPoints = 1
				} else {
					sumOfCardPoints = sumOfCardPoints * 2
				}
			}
		}
		sumOfPoints += sumOfCardPoints
	})

	return sumOfPoints

}

func parseLine(line string) ([]int, []int) {
	splitByColon := strings.Split(line, ":")
	if len(splitByColon) != 2 {
		panic("invalid format")
	}

	splitByPipe := strings.Split(splitByColon[1], "|")
	if len(splitByPipe) != 2 {
		panic("invalid format")
	}

	winningStrings := strings.Fields(splitByPipe[0])
	winningNumbers := make([]int, len(winningStrings))
	for i, numStr := range winningStrings {
		var err error
		winningNumbers[i], err = strconv.Atoi(numStr)
		util.PanicIfErr(err)
	}

	elfsStrings := strings.Fields(splitByPipe[1])
	elfsNumbers := make([]int, len(elfsStrings))
	for i, numStr := range elfsStrings {
		var err error
		elfsNumbers[i], err = strconv.Atoi(numStr)
		util.PanicIfErr(err)
	}

	return winningNumbers, elfsNumbers

}
