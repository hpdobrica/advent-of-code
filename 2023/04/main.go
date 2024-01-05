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

	sumOfPoints, totalCardsProcessed := processInput(file)

	fmt.Println("Total point worth of cards is ", sumOfPoints)
	fmt.Println("Total cards processed with duplicates is ", totalCardsProcessed)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)

}

func processInput(file *os.File) (int, int) {
	sumOfPoints := 0
	totalCardsProcessed := 0
	duplicatesMap := make(map[int]int)
	util.ForLineOfFile(file, func(line string) {
		scratchcard := parseScratchcard(line)

		numberOfHits := processScratchcard(scratchcard, duplicatesMap)
		totalCardsProcessed++
		cardPoints := numberOfHitsToCardPoints(numberOfHits)

		numOfDuplicates, exists := duplicatesMap[scratchcard.CardNumber]
		if exists {
			for i := 0; i < numOfDuplicates; i++ {
				processScratchcard(scratchcard, duplicatesMap)
				totalCardsProcessed++
			}
		}

		sumOfPoints += cardPoints
	})

	return sumOfPoints, totalCardsProcessed

}

func processScratchcard(scratchcard Scratchcard, duplicatesMap map[int]int) int {
	sumOfCardPoints := 0
	winningMap := make(map[int]bool)

	for _, winningNum := range scratchcard.WinningNumbers {
		_, exists := winningMap[winningNum]
		if !exists {
			winningMap[winningNum] = true
		}
	}

	numberOfHits := 0
	for _, elfsNum := range scratchcard.YourNumbers {
		_, exists := winningMap[elfsNum]
		if exists {
			numberOfHits++

			if sumOfCardPoints == 0 {
				sumOfCardPoints = 1
			} else {
				sumOfCardPoints = sumOfCardPoints * 2
			}
		}
	}

	for i := 1; i <= numberOfHits; i++ {
		oldNum, exists := duplicatesMap[scratchcard.CardNumber+i]
		if exists {
			duplicatesMap[scratchcard.CardNumber+i] = oldNum + 1
		} else {
			duplicatesMap[scratchcard.CardNumber+i] = 1
		}
	}

	return numberOfHits
}

type Scratchcard struct {
	CardNumber     int
	WinningNumbers []int
	YourNumbers    []int
}

func parseScratchcard(line string) Scratchcard {
	splitByColon := strings.Split(line, ":")
	if len(splitByColon) != 2 {
		panic("invalid format")
	}

	cardNumberString := strings.Fields(splitByColon[0])[1]
	cardNumber, err := strconv.Atoi(cardNumberString)
	if err != nil {
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
	yourNumbers := make([]int, len(elfsStrings))
	for i, numStr := range elfsStrings {
		var err error
		yourNumbers[i], err = strconv.Atoi(numStr)
		util.PanicIfErr(err)
	}

	return Scratchcard{
		CardNumber:     cardNumber,
		WinningNumbers: winningNumbers,
		YourNumbers:    yourNumbers,
	}

}

func numberOfHitsToCardPoints(numberOfHits int) int {
	if numberOfHits == 0 {
		return 0
	}
	score := 1
	for i := 1; i < numberOfHits; i++ {
		score = score * 2
	}
	return score

}
