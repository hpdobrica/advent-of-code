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

	safe, actuallySafe := processInput(file)

	fmt.Println("number of safe reports ", safe)
	fmt.Println("number of actually safe reports ", actuallySafe)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)
}

func processInput(file *os.File) (int, int) {

	safeReports := 0
	actuallySafeReports := 0

	util.ForLineOfFile(file, func(line string) {
		reportStrs := strings.Split(line, " ")
		report := make([]int, len(reportStrs))
		for i, str := range reportStrs {
			report[i], _ = strconv.Atoi(str)
		}

		// if safe, _ := isReportSafe(report); safe {
		// 	safeReports += 1
		// }
		if isReportActuallySafe(report) {
			actuallySafeReports += 1
		}

	})
	return safeReports, actuallySafeReports

}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isReportActuallySafe(report []int) bool {
	fmt.Println("<<<<< starting to process report====", report)
	safe, mistake := isReportSafe(report)
	defer fmt.Println(">>>>> finished processing report====", report, safe)

	if safe {
		return safe
	}

	for i := mistake; i > mistake-3; i-- {
		if i < 0 {
			continue
		}
		newReport := remove(report, i)
		fmt.Println("x x x x x x x trying a report with removed", i, newReport, report)
		safe, _ = isReportSafe(newReport)
		if safe {
			return safe
		}
	}

	return false

}

func isReportSafe(report []int) (bool, int) {
	fmt.Println(report)
	direction := ""
	for i := range report {
		if i == 0 {
			continue
		}
		biggerNum := 0
		smallerNum := 0

		if report[i-1] > report[i] {
			if direction != "DESC" && direction != "" {
				fmt.Println("direction changing to desc at", i)
				return false, i
			}
			direction = "DESC"
			biggerNum = report[i-1]
			smallerNum = report[i]
		}

		if report[i-1] < report[i] {
			if direction != "ASC" && direction != "" {
				fmt.Println("direction changing to asc at", i)
				return false, i
			}
			direction = "ASC"
			biggerNum = report[i]
			smallerNum = report[i-1]
		}

		if report[i-1] == report[i] {
			fmt.Println("numbers are equal at", i)
			return false, i
		}

		if biggerNum-smallerNum > 3 {
			fmt.Println("delta too big at", i)
			return false, i
		}

	}

	return true, -1

}
