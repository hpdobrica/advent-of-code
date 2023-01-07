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
	stacks9000 []Stack
	stacks9001 []Stack
	stage      Stage
}

type Stack []string

func (d *Dock) loadDock(input []string) {

	// calculate the number of stacks
	for index, line := range input {
		if index == 0 {
			re := regexp.MustCompile(`.*(\d) `)
			match := re.FindStringSubmatch(line)
			numOfStacks, _ := strconv.Atoi(match[1])
			for i := 0; i < numOfStacks; i++ {
				d.stacks9000 = append(d.stacks9000, Stack{})
				d.stacks9001 = append(d.stacks9001, Stack{})
			}
			continue
		}

		// populate each stack
		re := regexp.MustCompile(`.(.). .(.). .(.). .(.). .(.). .(.). .(.). .(.). .(.).`)
		match := re.FindStringSubmatch(line)
		for index, crate := range match[1:] {
			if crate != " " {
				d.stacks9000[index] = append(d.stacks9000[index], crate)
				d.stacks9001[index] = append(d.stacks9001[index], crate)
			}
		}

	}
	fmt.Println("Docks loaded, printing docks...")
	d.printDocks()
}

func (d *Dock) moveCrates9000(amount, source, target int) {

	for i := 0; i < amount; i++ {
		crate := pop(&d.stacks9000[source-1])
		d.stacks9000[target-1] = append(d.stacks9000[target-1], crate)
	}

}

func (d *Dock) moveCrates9001(amount, source, target int) {
	crates := multiPop(&d.stacks9001[source-1], amount)
	for _, crate := range crates {
		d.stacks9001[target-1] = append(d.stacks9001[target-1], crate)
	}
}

func (d *Dock) printDocks() {
	fmt.Println("stack of CrateMover 9000:")
	for _, stack := range d.stacks9000 {
		fmt.Println(stack)
	}
	fmt.Println("stack of CrateMover 9001:")
	for _, stack := range d.stacks9001 {
		fmt.Println(stack)
	}
}

func (d *Dock) printTopOfDock() {
	fmt.Println("Top of dock of StackMover 9000:")
	for _, stack := range d.stacks9000 {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Println()
	fmt.Println("Top of dock of StackMover 9001:")
	for _, stack := range d.stacks9001 {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Println()
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

			d.moveCrates9000(amount, source, target)
			d.moveCrates9001(amount, source, target)
		}

	})

	fmt.Println("processinc completed, printing docks...")

	d.printDocks()
	d.printTopOfDock()

}

func NewDock() *Dock {
	dock := &Dock{}
	dock.stage = Load
	dock.stacks9000 = []Stack{}
	dock.stacks9001 = []Stack{}

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

func multiPop(s *Stack, amount int) []string {
	len := len(*s)
	popped := (*s)[len-amount:]
	*s = (*s)[:len-amount]
	return popped
}
