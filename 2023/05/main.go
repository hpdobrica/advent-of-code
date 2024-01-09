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

	lowestLocationNumber := processInput(file)

	fmt.Println("Lowest location number is ", lowestLocationNumber)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)
}

type Mapping struct {
	sourceName      string
	destinationName string
	ranges          []MappingRange
}

type MappingRange struct {
	sourceStart      int
	destinationStart int
	length           int
}

func processInput(file *os.File) int {

	var currentIndices []int
	mappings := make([]Mapping, 0)

	var currentMapping Mapping

	util.ForLineOfFile(file, func(line string) {
		if len(line) > 0 {
			if line[0:len("seeds")] == "seeds" {
				currentIndices = parseSeedLine(line)
			}

			if line[len(line)-1:] == ":" {
				currentMapping = parseMappingHeader(line)
			}

			if line[0] >= '0' && line[0] <= '9' {
				currentMapping.ranges = append(currentMapping.ranges, parseMappingRange(line))
			}
		}

		if len(line) == 0 && currentMapping.sourceName != "" {
			currentIndices = processMapping(currentMapping, currentIndices)

			// write down? maybe wont be needed TODO
			mappings = append(mappings, currentMapping)
			// reset
			currentMapping = Mapping{}
		}

	})

	currentIndices = processMapping(currentMapping, currentIndices)

	fmt.Println("indices", currentIndices)

	return findMin(currentIndices)
}

func findMin(arr []int) int {
	min := arr[0]

	for _, el := range arr {
		if el < min {
			min = el
		}
	}
	return min
}

func processMapping(mapping Mapping, targets []int) []int {
	newTargets := make([]int, len(targets))
	for i, target := range targets {
		for _, r := range mapping.ranges {
			if target >= r.sourceStart && target <= r.sourceStart+r.length {

				diff := target - r.sourceStart
				newTargets[i] = r.destinationStart + diff
				break
			}
		}
		if newTargets[i] == 0 {
			fmt.Println("defaulting mapping for ", target)
			newTargets[i] = target
		}
	}
	return newTargets
}

func parseSeedLine(line string) []int {
	seeds := make([]int, 0)

	splitByColon := strings.Split(line, ":")
	if len(splitByColon) != 2 {
		panic("seeds format invalid, split by colon not len 2")
	}
	seedStrings := strings.Split(strings.TrimSpace(splitByColon[1]), " ")
	for _, seedString := range seedStrings {
		seed, err := strconv.Atoi(seedString)
		util.PanicIfErr(err)
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseMappingHeader(line string) Mapping {
	splitBySpace := strings.Split(line, " ")
	if len(splitBySpace) != 2 {
		panic("mapping header format invalid - split by space not len 2")
	}
	splitByDash := strings.Split(splitBySpace[0], "-")
	if len(splitByDash) != 3 {
		panic("mapping header format invalid - split by dash not len 3")
	}
	return Mapping{
		sourceName:      splitByDash[0],
		destinationName: splitByDash[2],
		ranges:          make([]MappingRange, 0),
	}
}

func parseMappingRange(line string) MappingRange {
	splitBySpace := strings.Split(line, " ")
	if len(splitBySpace) != 3 {
		panic("mapping format invalid")
	}

	sourceStart, err := strconv.Atoi(splitBySpace[1])
	util.PanicIfErr(err)

	destinationStart, err := strconv.Atoi(splitBySpace[0])
	util.PanicIfErr(err)

	length, err := strconv.Atoi(splitBySpace[2])
	util.PanicIfErr(err)

	return MappingRange{
		sourceStart:      sourceStart,
		destinationStart: destinationStart,
		length:           length,
	}
}
