package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	check(err)
	defer inputFile.Close()

	index := findSignalStartIndex(inputFile)

	fmt.Println("Sequence start index is", index+1)

}

func findSignalStartIndex(inputFile *os.File) int {
	sequenceIndex := 0
	fourLetterSequence := ""
	forCharOfFile(inputFile, func(i int, c string) error {
		if len(fourLetterSequence) == 4 {
			fourLetterSequence = fourLetterSequence[1:]
		}
		fourLetterSequence += c
		if isStartSequence(fourLetterSequence) {
			sequenceIndex = i
			return io.EOF
		}

		return nil

	})
	return sequenceIndex
}

func isStartSequence(s string) bool {
	if len(s) < 4 {
		return false
	}
	uniques := make(map[string]bool)

	for _, r := range s {
		c := string(r)
		if _, val := uniques[c]; !val {
			uniques[c] = true
		} else {
			return false
		}
	}
	return true

}

// general utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func forCharOfFile(file *os.File, fn func(int, string) error) {
	reader := bufio.NewReader(file)

	for i := 0; ; i++ {
		rune, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		if err := fn(i, string(rune)); err == io.EOF {
			break
		}
	}

}

func forLineOfFile(file *os.File, fn func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
}
