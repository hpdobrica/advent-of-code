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

	sumOfEngineParts, sumOfGearRatios := getSumOfEngineParts(engineMap)
	fmt.Println("Sum of engine parts is ", sumOfEngineParts)
	fmt.Println("Sum of gear ratios is ", sumOfGearRatios)
	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)

}

func getSumOfEngineParts(engineMap EngineMap) (int, int) {
	sumOfEngineParts := 0
	sumOfGearRatios := 0

	tmpNumString := ""
	tmpNumICol := 0

	shouldProcessNumber := false

	gearMap := make(map[string][]int)

	for iRow, row := range engineMap {
		for iCol, col := range row {
			if col.kind == KIND_NUMBER {
				if tmpNumICol == 0 && tmpNumString == "" {
					tmpNumICol = iCol
				}
				tmpNumString += col.value

				if iCol == len(row)-1 {
					shouldProcessNumber = true
				}
			} else {
				if tmpNumICol == 0 && tmpNumString == "" {
					continue
				}

				shouldProcessNumber = true
			}
			if shouldProcessNumber {
				partSymbols := getPartSymbols(engineMap, tmpNumString, iRow, tmpNumICol)
				if len(partSymbols) != 0 {
					numInt, err := strconv.Atoi(tmpNumString)
					util.PanicIfErr(err)
					sumOfEngineParts += numInt

					for _, symbol := range partSymbols {
						key := fmt.Sprintf("%d-%d", symbol.row, symbol.col)
						if _, exists := gearMap[key]; !exists {
							gearMap[key] = []int{numInt}
						} else {
							gearMap[key] = append(gearMap[key], numInt)
						}
					}
				}
				tmpNumICol = 0
				tmpNumString = ""
				shouldProcessNumber = false
			}
		}
	}

	for _, gearComponents := range gearMap {
		if len(gearComponents) == 2 {
			sumOfGearRatios += gearComponents[0] * gearComponents[1]
		}
	}

	return sumOfEngineParts, sumOfGearRatios
}

func getPartSymbols(engineMap EngineMap, numString string, iRow, iCol int) []EngineSymbol {
	symbols := make([]EngineSymbol, 0)

	// check left
	if iCol > 0 && engineMap[iRow][iCol-1].kind == KIND_SYMBOL {
		symbols = append(symbols, EngineSymbol{
			row:   iRow,
			col:   iCol - 1,
			value: engineMap[iRow][iCol-1].value,
		})
	}

	// check right
	if iCol+len(numString) < len(engineMap[iRow])-1 && engineMap[iRow][iCol+len(numString)].kind == KIND_SYMBOL {
		symbols = append(symbols, EngineSymbol{
			row:   iRow,
			col:   iCol + len(numString),
			value: engineMap[iRow][iCol+len(numString)].value,
		})
	}

	// check above
	if iRow > 0 {
		// check above left diagonal
		if iCol > 0 && engineMap[iRow-1][iCol-1].kind == KIND_SYMBOL {
			symbols = append(symbols, EngineSymbol{
				row:   iRow - 1,
				col:   iCol - 1,
				value: engineMap[iRow-1][iCol-1].value,
			})
		}

		// check above right diagonal
		if iCol+len(numString) < len(engineMap[iRow])-1 && engineMap[iRow-1][iCol+len(numString)].kind == KIND_SYMBOL {
			symbols = append(symbols, EngineSymbol{
				row:   iRow - 1,
				col:   iCol + len(numString),
				value: engineMap[iRow-1][iCol+len(numString)].value,
			})
		}

		// check directly above
		for i := range numString {
			if engineMap[iRow-1][iCol+i].kind == KIND_SYMBOL {
				symbols = append(symbols, EngineSymbol{
					row:   iRow - 1,
					col:   iCol + i,
					value: engineMap[iRow-1][iCol+i].value,
				})
			}
		}

	}

	if iRow < len(engineMap)-1 {
		// check below left diagonal
		if iCol > 0 && engineMap[iRow+1][iCol-1].kind == KIND_SYMBOL {
			symbols = append(symbols, EngineSymbol{
				row:   iRow + 1,
				col:   iCol - 1,
				value: engineMap[iRow+1][iCol-1].value,
			})
		}

		// check below right diagonal
		if iCol+len(numString) < len(engineMap[iRow])-1 && engineMap[iRow+1][iCol+len(numString)].kind == KIND_SYMBOL {
			symbols = append(symbols, EngineSymbol{
				row:   iRow + 1,
				col:   iCol + len(numString),
				value: engineMap[iRow+1][iCol+len(numString)].value,
			})
		}

		// check directly below
		for i := range numString {
			if engineMap[iRow+1][iCol+i].kind == KIND_SYMBOL {
				symbols = append(symbols, EngineSymbol{
					row:   iRow + 1,
					col:   iCol + i,
					value: engineMap[iRow+1][iCol+i].value,
				})
			}
		}
	}

	return symbols
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

type EngineSymbol struct {
	row   int
	col   int
	value string
}

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
