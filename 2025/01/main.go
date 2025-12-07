package main

import (
	"fmt"
	"os"
	"strconv"
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

	password := processInput(file)

	fmt.Println("total times dial at zero (password): ", password)

	elapsed := time.Since(start)
	fmt.Println("Done after ", elapsed)

}

func processInput(file *os.File) int {

	dial := NewDial()
	numberOfZeros := 0
	util.ForLineOfFile(file, func(command string) {
		positionAfterMove := dial.move(command)
		if positionAfterMove == 0 {
			numberOfZeros += 1
		}
	})

	return numberOfZeros
}

type Dial struct {
	dial     []int
	position int
}

func NewDial() Dial {
	dial := Dial{}
	dial.dial = make([]int, 100)
	for i := range dial.dial {
		dial.dial[i] = i
	}

	dial.position = 50

	return dial
}

func (d *Dial) move(command string) int {
	var direction string = string(command[0])
	distance, _ := strconv.Atoi(command[1:])
	fmt.Println(direction, distance)

	if direction == "R" {
		d.position = (d.position + distance) % 100
	}
	if direction == "L" {
		newPosition := d.position - distance
		for newPosition < 0 {
			newPosition = 100 + newPosition
		}
		d.position = newPosition

	}
	fmt.Println("lands at", d.position)
	return d.position
}
