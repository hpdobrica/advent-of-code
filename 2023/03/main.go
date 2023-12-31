package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {

	start := time.Now()

	file, err := os.Open("./simple-input.txt")
	util.PanicIfErr(err)

	util.ForLineOfFile(file, func(line string) {

	})

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)

}
