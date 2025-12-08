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
		_, zerosEncounteredDuringRotation := dial.move(command)
		newNumberOfZeros := 0
		newNumberOfZeros += zerosEncounteredDuringRotation

		fmt.Printf(" +%d\n", newNumberOfZeros)
		numberOfZeros += newNumberOfZeros
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

func (d *Dial) move(command string) (int, int) {
	var zerosEncounteredDuringRotation int = 0
	var direction string = string(command[0])
	distance, _ := strconv.Atoi(command[1:])

	if direction == "R" {
		newPosition := (d.position + distance) % 100
		zerosEncounteredDuringRotation += (d.position + distance) / 100
		d.position = newPosition

	}
	if direction == "L" {
		newPosition := d.position - distance
		if newPosition == 0 {
			zerosEncounteredDuringRotation += 1
		}
		for newPosition < 0 {
			newPosition = 100 + newPosition
			zerosEncounteredDuringRotation += 1
			if newPosition == 0 {
				zerosEncounteredDuringRotation += 1
			}

		}
		if d.position == 0 {
			zerosEncounteredDuringRotation -= 1
		}

		d.position = newPosition

	}
	fmt.Printf("%s (%d)", command, d.position)
	return d.position, zerosEncounteredDuringRotation
}
