package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {

	start := time.Now()

	file, err := os.Open("./input.txt")
	util.PanicIfErr(err)

	engineMap := fileToEngineMap(file)
	fmt.Println(engineMap)

	sumOfEngineParts := getSumOfEngineParts(engineMap)
	fmt.Println("Sum of engine parts is ", sumOfEngineParts)
	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)

}

func getSumOfEngineParts(engineMap EngineMap) int {
	sum := 0

	tmpNumString := ""
	tmpNumICol := 0

	for iRow, row := range engineMap {
		for iCol, col := range row {
			if col.kind == KIND_NUMBER {
				if tmpNumICol == 0 && tmpNumString == "" {
					tmpNumICol = iCol
				}
				tmpNumString += col.value

				if iCol == len(row)-1 {
					sum += getPartNumberOrZero(engineMap, tmpNumString, iRow, tmpNumICol)
					tmpNumICol = 0
					tmpNumString = ""
				}
			} else if col.kind == KIND_SYMBOL {
				if tmpNumICol == 0 && tmpNumString == "" {
					continue
				}

				numInt, err := strconv.Atoi(tmpNumString)
				util.PanicIfErr(err)

				sum += numInt

				tmpNumICol = 0
				tmpNumString = ""

			} else if col.kind == KIND_EMPTY {
				if tmpNumICol == 0 && tmpNumString == "" {
					continue
				}

				sum += getPartNumberOrZero(engineMap, tmpNumString, iRow, tmpNumICol)
				tmpNumICol = 0
				tmpNumString = ""
			}
		}
	}

	return sum
}

func getPartNumberOrZero(engineMap EngineMap, numString string, iRow, iCol int) int {
	numInt, err := strconv.Atoi(numString)
	util.PanicIfErr(err)
	partNumber := numInt

	// check left
	if iCol > 0 && engineMap[iRow][iCol-1].kind == KIND_SYMBOL {
		return partNumber
	}

	// check above
	if iRow > 0 {
		// check above left diagonal
		if iCol > 0 && engineMap[iRow-1][iCol-1].kind == KIND_SYMBOL {
			return partNumber
		}

		// check above right diagonal
		if iCol+len(numString) < len(engineMap[iRow])-1 && engineMap[iRow-1][iCol+len(numString)].kind == KIND_SYMBOL {
			return partNumber
		}

		// check directly above
		for i := range numString {
			if engineMap[iRow-1][iCol+i].kind == KIND_SYMBOL {
				return partNumber
			}
		}

	}

	if iRow < len(engineMap)-1 {
		// check below left diagonal
		if iCol > 0 && engineMap[iRow+1][iCol-1].kind == KIND_SYMBOL {
			return partNumber
		}

		// check below right diagonal
		if iCol+len(numString) < len(engineMap[iRow])-1 && engineMap[iRow+1][iCol+len(numString)].kind == KIND_SYMBOL {
			return partNumber
		}

		// check directly below
		for i := range numString {
			if engineMap[iRow+1][iCol+i].kind == KIND_SYMBOL {
				return partNumber
			}
		}
	}

	return 0
}

type EngineMap [][]EngineChar
type EngineChar struct {
	value string
	kind  EngineCharKind
}

type EngineCharKind int

const (
	KIND_EMPTY EngineCharKind = iota
	KIND_NUMBER
	KIND_SYMBOL
	KIND_UNKNOWN
)

func fileToEngineMap(file *os.File) EngineMap {

	engineMap := make(EngineMap, 0)
	util.ForLineOfFile(file, func(line string) {
		engineRow := make([]EngineChar, len(line))

		for i, char := range line {
			kind := KIND_UNKNOWN
			if char == '.' {
				kind = KIND_EMPTY
			} else if unicode.IsDigit(char) {
				kind = KIND_NUMBER
			} else {
				kind = KIND_SYMBOL
			}

			if kind == KIND_UNKNOWN {
				panic("character cant be unknown type!")
			}

			engineRow[i] = EngineChar{
				value: string(char),
				kind:  kind,
			}
		}

		engineMap = append(engineMap, engineRow)
	})

	return engineMap
}
