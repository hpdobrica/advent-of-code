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

	sum := 0

	util.ForLineOfFile(file, func(line string) {
		sum += calculateLineDigit(line)
	})
	fmt.Println("result: ", sum)
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
