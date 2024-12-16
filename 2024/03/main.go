package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")
	util.PanicIfErr(err)

	sum := processInput(file)

	fmt.Println("total sum is ", sum)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)
}

func processInput(file *os.File) int {
	regex := regexp.MustCompile(`mul\((\d*),(\d*)\)|do\(\)|don't\(\)`)
	result := 0
	shouldWork := true

	util.ForLineOfFile(file, func(line string) {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fullStr, oneStr, twoStr := match[0], match[1], match[2]
			if fullStr == "don't()" {
				shouldWork = false
				continue
			} else if fullStr == "do()" {
				shouldWork = true
				continue
			}
			if !shouldWork {
				continue
			}
			one, _ := strconv.Atoi(oneStr)
			two, _ := strconv.Atoi(twoStr)
			result += one * two
		}
	})
	return result

}
