package util

import (
	"bufio"
	"os"
)

func ForLineOfFile(file *os.File, fn func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
}

func PanicIfErr(e error) {
	if e != nil {
		panic(e)
	}
}
