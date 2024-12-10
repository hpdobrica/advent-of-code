package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/hpdobrica/advent-of-code/util"
)

func Insert[T cmp.Ordered](ts []T, t T) []T {
	i, _ := slices.BinarySearch(ts, t) // find slot
	return slices.Insert(ts, i, t)
}

func absDiff(x int, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	util.PanicIfErr(err)

	drift, similarity := processInput(file)

	fmt.Println("total drift is ", drift)
	fmt.Println("total silmilarity is ", similarity)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)
}

func processInput(file *os.File) (int, int) {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	util.ForLineOfFile(file, func(line string) {
		numStrings := strings.Fields(line)

		leftNum, err := strconv.Atoi(numStrings[0])
		if err != nil {
			util.PanicIfErr(err)
		}
		rightNum, err := strconv.Atoi(numStrings[1])
		if err != nil {
			util.PanicIfErr(err)
		}

		leftList = Insert(leftList, leftNum)
		rightList = Insert(rightList, rightNum)

	})

	// fmt.Println(leftList)
	// fmt.Println(rightList)

	totalDrift := 0

	for i := range leftList {
		totalDrift += absDiff(leftList[i], rightList[i])
	}

	totalSimilarity := 0

	rightCount := make(map[int]int)

	for _, n := range rightList {
		rightCount[n] += 1
	}

	for _, n := range leftList {
		totalSimilarity += n * rightCount[n]
	}

	return totalDrift, totalSimilarity

}
