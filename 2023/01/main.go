package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {
	file, err := os.Open("./input.txt")

	util.PanicIfErr(err)

	firstSum := 0
	secondSum := 0

	util.ForLineOfFile(file, func(line string) {
		firstSum += calculateLineDigit(line)
		secondSum += calculateLineDigitWithLetters(line)
	})
	fmt.Println("sum without letters converted: ", firstSum)
	fmt.Println("sum with letters converted", secondSum)
}

func calculateLineDigitWithLetters(line string) int {

	type numberDetails struct {
		Value int
		Index int
	}

	dict := make(map[string]numberDetails)

	dict["one"] = numberDetails{
		Value: 1,
		Index: 0,
	}
	dict["two"] = numberDetails{
		Value: 2,
		Index: 0,
	}
	dict["three"] = numberDetails{
		Value: 3,
		Index: 0,
	}
	dict["four"] = numberDetails{
		Value: 4,
		Index: 0,
	}
	dict["five"] = numberDetails{
		Value: 5,
		Index: 0,
	}
	dict["six"] = numberDetails{
		Value: 6,
		Index: 0,
	}
	dict["seven"] = numberDetails{
		Value: 7,
		Index: 0,
	}
	dict["eight"] = numberDetails{
		Value: 8,
		Index: 0,
	}
	dict["nine"] = numberDetails{
		Value: 9,
		Index: 0,
	}

	result := 0
	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(string(line[i]))
		if err == nil {
			result += num * 10
			break
		}
		foundWord := false
		for name, details := range dict {
			if line[i] == name[details.Index] {
				if details.Index == len(name)-1 {
					result += details.Value * 10
					foundWord = true
					break
				}
				if newDetail, ok := dict[name]; ok {
					newDetail.Index++
					dict[name] = newDetail
				}

			}
		}
		if foundWord {
			break
		}

	}

	for name := range dict {
		if newDetail, ok := dict[name]; ok {
			newDetail.Index = len(name) - 1
			dict[name] = newDetail
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(line[i]))
		if err == nil {
			result += num
			break
		}
		foundWord := false
		for name, details := range dict {
			if line[i] == name[details.Index] {
				if details.Index == 0 {
					result += details.Value
					foundWord = true
					break
				}
				if newDetail, ok := dict[name]; ok {
					newDetail.Index--
					dict[name] = newDetail
				}
			}
		}
		if foundWord {
			break
		}
	}

	return result
}

func calculateLineDigit(line string) int {
	result := 0
	for i := 0; i < len(line); i++ {
		num, err := strconv.Atoi(string(line[i]))
		if err == nil {
			result += num * 10
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(line[i]))
		if err == nil {
			result += num
			break
		}
	}

	return result

}
