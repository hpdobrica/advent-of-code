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

	useSimpleInput := false

	var fileName string
	if useSimpleInput {
		fileName = "./simple-input.txt"
	} else {
		fileName = "./input.txt"
	}
	file, err := os.Open(fileName)
	util.PanicIfErr(err)

	sum, badSum := processInput(file)

	fmt.Println("total xmases found ", sum)

	fmt.Println("total fixed xmases found ", badSum)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)
}

func processInput(file *os.File) (int, int) {

	rules := make(map[int]map[string][]int)
	updates := make([][]int, 0)
	parsingRules := true
	util.ForLineOfFile(file, func(line string) {
		if line == "" {
			fmt.Println("ONE BLANK")
			parsingRules = false
			return
		}
		if parsingRules {
			ruleStrings := strings.Split(line, "|")
			left, _ := strconv.Atoi(ruleStrings[0])
			right, _ := strconv.Atoi(ruleStrings[1])

			if _, ok := rules[left]; !ok {
				rules[left] = make(map[string][]int)
			}
			if _, ok := rules[right]; !ok {
				rules[right] = make(map[string][]int)
			}

			appendRule(rules[left], "before", right)
			appendRule(rules[right], "after", left)

		} else {
			currentUpdatesStr := strings.Split(line, ",")
			currentUpdates := make([]int, len(currentUpdatesStr))
			for i, s := range currentUpdatesStr {
				n, _ := strconv.Atoi(s)
				currentUpdates[i] = n
			}
			updates = append(updates, currentUpdates)
		}
	})

	fmt.Println(rules)

	sum := 0
	badSum := 0

	for _, updateSeq := range updates {
		updateSeqValid, swapI, swapJ := checkSeq(updateSeq, rules)

		// fmt.Println("finally", updateSeq, updateSeqValid)
		if !updateSeqValid {
			for updateSeqValid == false {
				tmp := updateSeq[swapJ]
				updateSeq[swapJ] = updateSeq[swapI]
				updateSeq[swapI] = tmp
				updateSeqValid, swapI, swapJ = checkSeq(updateSeq, rules)
			}
			badSum += updateSeq[len(updateSeq)/2]
		} else {
			sum += updateSeq[len(updateSeq)/2]
		}

	}

	return sum, badSum

}

func checkSeq(updateSeq []int, rules map[int]map[string][]int) (bool, int, int) {
	for i, current := range updateSeq {

		before := updateSeq[:i]
		after := updateSeq[i+1:]

		beforeClear := false
		afterClear := false
		// fmt.Println("working on", current, before, after)
		if beforeRule, ok := rules[current]["after"]; ok {

			for j, u := range before {
				tmpBeforeClear := false
				for _, r := range beforeRule {
					// fmt.Println("checking before", before, u, r)
					if u == r {
						tmpBeforeClear = true
					}
				}
				beforeClear = tmpBeforeClear
				// fmt.Println("before clear", beforeClear)
				if !beforeClear {
					fmt.Println("failing on", updateSeq, current, i, j)
					return false, i, j
				}
			}
			if len(before) == 0 {
				beforeClear = true
			}

		} else {
			beforeClear = true
		}
		if afterRule, ok := rules[current]["before"]; ok {
			for j, u := range after {
				tmpAfterClear := false
				for _, r := range afterRule {
					// fmt.Println("checking after", after, u, r)
					if u == r {
						tmpAfterClear = true
					}
				}
				afterClear = tmpAfterClear
				// fmt.Println("after clear", afterClear)
				if !afterClear {
					fmt.Println("failing on", updateSeq, current, i, j+i+1, afterRule)
					return false, i, j + i + 1
				}
			}
			if len(after) == 0 {
				afterClear = true
			}
		} else {
			afterClear = true
		}

		// fmt.Println(updateSeq)
		// fmt.Println(current, beforeClear, afterClear)

	}
	// fmt.Println("finally", updateSeq, updateSeqValid)
	return true, -1, -1

}

func appendRule(inputMap map[string][]int, key string, numToAdd int) {
	if _, ok := inputMap[key]; ok {

		inputMap[key] = append(inputMap[key], numToAdd)
	} else {
		inputMap[key] = make([]int, 1)
		inputMap[key][0] = numToAdd
	}
}
