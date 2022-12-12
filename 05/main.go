package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stage int

const (
	Load Stage = iota
	Move
)

type Dock struct {
	stacks []Stack
	stage  Stage
}

type Stack []string

func (d *Dock) loadDock(input []string) {

	// numOfStacks := 0
	for index, line := range input {
		// fmt.Println(line, index)
		if index == 0 {
			re := regexp.MustCompile(`.*(\d) `)
			match := re.FindStringSubmatch(line)
			// fmt.Println(match)
			numOfStacks, _ := strconv.Atoi(match[1])
			for i := 0; i < numOfStacks; i++ {
				d.stacks = append(d.stacks, Stack{})
			}
			continue
		}

		re := regexp.MustCompile(`.(.). .(.). .(.). .(.). .(.). .(.). .(.). .(.). .(.).`)
		match := re.FindStringSubmatch(line)
		for index, crate := range match[1:] {
			if crate != " " {
				d.stacks[index] = append(d.stacks[index], crate)
			}
		}

	}
}

func (d *Dock) moveCrates(amount, source, target int) {

	for i := 0; i < amount; i++ {
		crate := pop(&d.stacks[source-1])
		d.stacks[target-1] = append(d.stacks[target-1], crate)
	}

}

func (d *Dock) printDock() {
	for _, stack := range d.stacks {
		fmt.Println(stack)
	}
}

func (d *Dock) printTopOfDock() {
	fmt.Println("Top of dock:")
	for _, stack := range d.stacks {
		fmt.Print(stack[len(stack)-1])
	}
}

func (d *Dock) processInput(file *os.File) {
	dockInput := []string{}

	forLineOfFile(file, func(line string) {
		if d.stage == Load {
			if line == "" {
				reverse(dockInput)
				d.loadDock(dockInput)
				d.stage = Move
				return
			}
			dockInput = append(dockInput, line)
		}
		if d.stage == Move {
			re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
			match := re.FindStringSubmatch(line)
			amount, _ := strconv.Atoi(match[1])
			source, _ := strconv.Atoi(match[2])
			target, _ := strconv.Atoi(match[3])

			d.moveCrates(amount, source, target)
		}

	})

	d.printDock()
	d.printTopOfDock()

}

func NewDock() *Dock {
	dock := &Dock{}
	dock.stage = Load
	dock.stacks = []Stack{}

	return dock
}

func main() {
	inputFile, err := os.Open("input.txt")
	check(err)
	defer inputFile.Close()

	dock := NewDock()

	dock.processInput(inputFile)

}

// general utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func forLineOfFile(file *os.File, fn func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fn(line)
	}
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func pop(s *Stack) string {
	len := len(*s)
	last := (*s)[len-1]
	*s = (*s)[:len-1]
	return last
}
