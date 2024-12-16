package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {
	start := time.Now()

	useSimpleInput := false

	var fileName string
	if useSimpleInput {
		fileName = "./simple-input.txt"
	} else {
		fileName = "./input.txt"
	}
	file, err := os.Open(fileName)
	util.PanicIfErr(err)

	sum, xsum := processInput(file)

	fmt.Println("total xmases found ", sum)
	fmt.Println("total xmases Xs found ", xsum)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)
}

func processInput(file *os.File) (int, int) {

	textMatrix := []string{}

	util.ForLineOfFile(file, func(line string) {
		textMatrix = append(textMatrix, line)
	})

	matches := 0
	xMatches := 0

	for y, row := range textMatrix {
		for x, char := range row {
			if char == 'X' {
				// fmt.Println("found x", x, y)
				matches += checkDirections(textMatrix, x, y)
			}
			if char == 'A' {
				// fmt.Println("found a", x, y)
				xMatches += checkX(textMatrix, x, y)
			}
		}
	}
	return matches, xMatches

}

func checkDirections(matrix []string, x, y int) int {
	targetString := "MAS"

	directions := make(map[string][2]int)
	directions["up"] = [2]int{0, -1}
	directions["lup"] = [2]int{-1, -1}
	directions["rup"] = [2]int{1, -1}
	directions["down"] = [2]int{0, 1}
	directions["ldown"] = [2]int{-1, 1}
	directions["rdown"] = [2]int{1, 1}
	directions["left"] = [2]int{-1, 0}
	directions["right"] = [2]int{1, 0}

	matches := 0

	for _, vector := range directions {
		tmpX := x
		tmpY := y
		for i := range len(targetString) {
			tmpX = tmpX + vector[0]
			tmpY = tmpY + vector[1]
			// x += vector[0]
			// y += vector[1]
			if tmpX < 0 || tmpY < 0 {
				// fmt.Println("skipping", x, y)
				continue
			}
			// fmt.Println("y >=", tmpY >= len(matrix), tmpY, len(matrix))
			// fmt.Println("x >=", tmpX >= len(matrix[0]), tmpX, len(matrix[0]))
			if tmpY >= len(matrix) || tmpX >= len(matrix[0]) {
				fmt.Println("skipping", tmpX, tmpY)
				continue
			}

			// fmt.Printf("%b: %c != %c\n", matrix[tmpY][tmpX] != targetString[i], matrix[tmpY][tmpX], targetString[i])
			if matrix[tmpY][tmpX] != targetString[i] {
				// fmt.Println("breaking", tmpX, tmpY)
				break
			}
			fmt.Println("semi-match", i, len(targetString)-1)
			if i == len(targetString)-1 {
				// fmt.Println("full match", tmpX, tmpY)
				matches += 1
			}
		}
	}
	return matches
}

func checkX(matrix []string, x, y int) int {
	// targetString := "MAS"

	directions := make(map[string][2]int)
	// directions["up"] = [2]int{0, -1}
	directions["lup"] = [2]int{-1, -1}
	directions["rup"] = [2]int{1, -1}
	// directions["down"] = [2]int{0, 1}
	directions["ldown"] = [2]int{-1, 1}
	directions["rdown"] = [2]int{1, 1}
	// directions["left"] = [2]int{-1, 0}
	// directions["right"] = [2]int{1, 0}

	matches := 0
	lrDiagonal := ""
	rlDiagonal := ""

	for dir, vector := range directions {
		tmpX := x
		tmpY := y

		tmpX = tmpX + vector[0]
		tmpY = tmpY + vector[1]
		if tmpX < 0 || tmpY < 0 {
			// fmt.Println("skipping", x, y)
			continue
		}
		if tmpY >= len(matrix) || tmpX >= len(matrix[0]) {
			// fmt.Println("skipping", tmpX, tmpY)
			continue
		}

		if dir == "lup" || dir == "rdown" {
			lrDiagonal += string(matrix[tmpY][tmpX])
		}
		if dir == "rup" || dir == "ldown" {
			rlDiagonal += string(matrix[tmpY][tmpX])
		}

	}

	fmt.Println(lrDiagonal, rlDiagonal)
	if strings.Contains(lrDiagonal, "M") && strings.Contains(lrDiagonal, "S") && strings.Contains(rlDiagonal, "M") && strings.Contains(rlDiagonal, "S") {
		matches++
	}

	return matches
}
