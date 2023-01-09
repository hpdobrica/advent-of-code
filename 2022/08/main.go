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
	// f.getTreesVisibleFromTop(&visibleTrees)
	// f.getTreesVisibleFromRight(&visibleTrees)
	// f.getTreesVisibleFromBottom(&visibleTrees)
	// f.getTreesVisibleFromLeft(&visibleTrees)
	f.getTreesVisible("top", &visibleTrees)
	f.getTreesVisible("right", &visibleTrees)
	f.getTreesVisible("bottom", &visibleTrees)
	f.getTreesVisible("left", &visibleTrees)
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

func (f Forest) getTreesVisible(direction string, visibleTreesMap *map[string]bool) {
	iDir := ""
	jDir := ""
	iFirst := true
	if direction == "top" {
		iDir = "forward"
		jDir = "forward"
		iFirst = false
	} else if direction == "right" {
		iDir = "forward"
		jDir = "back"
		iFirst = true
	} else if direction == "bottom" {
		iDir = "back"
		jDir = "back"
		iFirst = false
	} else if direction == "left" {
		iDir = "forward"
		jDir = "forward"
		iFirst = true
	}

	directionalFor(f, iDir, func(i int, iVal []int) string {
		biggestSoFar := 0
		directionalFor(iVal, jDir, func(j int, jVal int) string {
			if getTree(f, i, j, iFirst) == 9 {
				(*visibleTreesMap)[makeTreeKey(i, j, iFirst)] = true
				return "break"
			}

			if jDir == "forward" && j == 0 || jDir == "back" && j == len(f[i])-1 {
				biggestSoFar = getTree(f, i, j, iFirst)
				(*visibleTreesMap)[makeTreeKey(i, j, iFirst)] = true
				return ""
			}

			if getTree(f, i, j, iFirst) > biggestSoFar {
				biggestSoFar = getTree(f, i, j, iFirst)
				(*visibleTreesMap)[makeTreeKey(i, j, iFirst)] = true
				return ""
			}
			return ""
		})
		return ""
	})
}

func getTree(f Forest, i, j int, iFirst bool) int {
	if iFirst {
		return f[i][j]
	} else {
		return f[j][i]
	}
}

func makeTreeKey(i, j int, iFirst bool) string {
	if iFirst {
		return strconv.Itoa(i) + "," + strconv.Itoa(j)
	} else {
		return strconv.Itoa(j) + "," + strconv.Itoa(i)
	}
}

func directionalFor[T any](arr []T, direction string, fn func(int, T) string) {
	if direction == "forward" {
		for i := 0; i < len(arr); i++ {
			if fn(i, arr[i]) == "break" {
				break
			}
		}
	}
	if direction == "back" {
		for i := len(arr) - 1; i >= 0; i-- {
			if fn(i, arr[i]) == "break" {
				break
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
