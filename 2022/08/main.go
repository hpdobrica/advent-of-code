package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Forest [][]int

func NewForest(inputFile *os.File) *Forest {

	f := Forest{}

	rowIndex := 0
	forLineOfFile(inputFile, func(line string) {
		f = append(f, []int{})

		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			f[rowIndex] = append(f[rowIndex], num)
		}
		rowIndex++
	})

	return &f
}

func (f Forest) Print() {
	for _, row := range f {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func (f Forest) PrintMask(trueMap map[string]bool) {
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"
	for i, row := range f {
		for j, col := range row {
			if trueMap[strconv.Itoa(i)+","+strconv.Itoa(j)] {
				fmt.Fprintf(os.Stdout, "%s%s", colorRed, strconv.Itoa(col))
				// fmt.Printf(strconv.Itoa(col))
			} else {
				fmt.Fprintf(os.Stdout, "%s%s", colorNone, strconv.Itoa(col))
				// fmt.Printf(" ")
			}

		}
		fmt.Println()
	}
}

func (f Forest) CountTreesVisibleFromEdge() int {
	startRow := 0
	endRow := len(f) - 1

	startCol := 0
	endCol := len(f[0]) - 1

	fmt.Println(startRow, endRow, startCol, endCol)
	fmt.Println(f)

	// visibleTrees := 0

	// for rowIndex, row := range f {
	// 	for colIndex, col := range row {
	// 		if rowIndex == startRow || rowIndex == endRow || colIndex == startCol || colIndex == endCol {
	// 			visibleTrees++
	// 			continue
	// 		}

	// 	}
	// }
	visibleTrees := map[string]bool{}
	f.getTreesVisibleFromTop(&visibleTrees)
	f.getTreesVisibleFromRight(&visibleTrees)
	f.getTreesVisibleFromBottom(&visibleTrees)
	f.getTreesVisibleFromLeft(&visibleTrees)
	fmt.Println(visibleTrees)
	count := countTruesInMap(visibleTrees)
	f.PrintMask(visibleTrees)

	return count
}

func countTruesInMap(m map[string]bool) int {
	count := 0
	for _, val := range m {
		if val {
			count++
		}
	}
	return count
}

func (f Forest) getTreesVisibleFromTop(visibleTreesMap *map[string]bool) {
	for i := 0; i < len(f); i++ {
		biggestSoFar := 0
		for j := 0; j < len(f[0]); j++ {
			if f[j][i] == 9 {
				(*visibleTreesMap)[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				break
			}
			if j == 0 {
				biggestSoFar = f[j][i]
				(*visibleTreesMap)[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				continue
			}
			if f[j][i] > biggestSoFar {
				biggestSoFar = f[j][i]
				(*visibleTreesMap)[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				continue
			}

		}
	}
	fmt.Println(*visibleTreesMap)
}

func (f Forest) getTreesVisibleFromRight(visibleTreesMap *map[string]bool) {
	for i := 0; i < len(f); i++ {
		biggestInRow := 0
		for j := len(f[i]) - 1; j >= 0; j-- {
			fmt.Println("processing", i, j)
			if f[i][j] == 9 {
				(*visibleTreesMap)[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				break
			}
			if j == len(f[i])-1 {
				fmt.Println("== in the first row, autoaccepted")
				biggestInRow = f[i][j]
				(*visibleTreesMap)[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				continue
			}
			if f[i][j] > biggestInRow {
				fmt.Println("== bigger than biggest, thus visible", f[i][j], ">", biggestInRow)
				biggestInRow = f[i][j]
				(*visibleTreesMap)[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				continue
			} else {
				fmt.Println("== not bigger than the biggest, thus continuing", f[i][j], "<=", biggestInRow)
				continue
			}

		}
	}
}

func (f Forest) getTreesVisibleFromBottom(visibleTreesMap *map[string]bool) {
	for i := len(f) - 1; i >= 0; i-- {
		biggestSoFar := 0
		for j := len(f[i]) - 1; j >= 0; j-- {
			fmt.Println("processing", j, i)
			if f[j][i] == 9 {
				(*visibleTreesMap)[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				break
			}
			if j == len(f)-1 {
				fmt.Println("== in the first row, autoaccepted")
				biggestSoFar = f[j][i]
				(*visibleTreesMap)[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				continue
			}
			if f[j][i] > biggestSoFar {
				fmt.Println("== bigger than the biggest so far", f[j][i], ">", biggestSoFar)
				biggestSoFar = f[j][i]
				(*visibleTreesMap)[strconv.Itoa(j)+","+strconv.Itoa(i)] = true
				continue
			} else {
				fmt.Println("== not bigger than the biggest so far, continuing", f[j][i], "<=", biggestSoFar)
				continue
			}

		}
	}
}

func (f Forest) getTreesVisibleFromLeft(visibleTreesMap *map[string]bool) {
	for i := 0; i < len(f); i++ {
		biggestSoFar := 0
		for j := 0; j < len(f[0]); j++ {
			if f[i][j] == 9 {
				(*visibleTreesMap)[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				break
			}
			if j == 0 {
				biggestSoFar = f[i][j]
				(*visibleTreesMap)[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				continue
			}
			if f[i][j] > biggestSoFar {
				biggestSoFar = f[i][j]
				(*visibleTreesMap)[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
				continue
			}

		}
	}
}

func main() {
	inputFile, err := os.Open("input.txt")
	check(err)

	forest := NewForest(inputFile)

	treesVisibleFromEdge := forest.CountTreesVisibleFromEdge()

	fmt.Println("there is total", treesVisibleFromEdge, "trees visible from edge")

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
