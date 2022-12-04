package main

import (
	"bufio"
	"fmt"
	"os"
)

type Move struct {
	name        string
	score       int
	winsAgainst *Move
}

var ROCK = Move{"ROCK", 1, nil}
var PAPER = Move{"PAPER", 2, nil}
var SCISSORS = Move{"SCISSORS", 3, nil}

const LOSE_SCORE = 0
const DRAW_SCORE = 3
const WIN_SCORE = 6

func main() {
	ROCK.winsAgainst = &SCISSORS
	PAPER.winsAgainst = &ROCK
	SCISSORS.winsAgainst = &PAPER

	file, err := os.Open("./input.txt")
	check(err)

	totalScore := calculateTotalScore(file)

	fmt.Println("Total score is", totalScore)

}

func calculateTotalScore(file *os.File) int {

	totalScore := 0

	forLineOfFile(file, func(line string) {
		totalScore += calculateScoreForRound(line[0:1], line[2:3])
	})

	return totalScore
}

func calculateScoreForRound(enemyMoveCode string, ourMoveCode string) int {

	roundScore := 0

	enemyMove, ourMove := moveDecode(enemyMoveCode, ourMoveCode)

	roundScore += ourMove.score

	if ourMove == enemyMove {
		roundScore += DRAW_SCORE
	} else if *ourMove.winsAgainst == enemyMove {
		roundScore += WIN_SCORE
	} else {
		roundScore += LOSE_SCORE
	}

	fmt.Println(enemyMove, ourMove, roundScore)

	return roundScore

}

func moveDecode(enemyMoveCode string, ourMoveCode string) (Move, Move) {
	var dictionary = map[string]Move{
		"A": ROCK,
		"X": ROCK,
		"B": PAPER,
		"Y": PAPER,
		"C": SCISSORS,
		"Z": SCISSORS,
	}

	ourMove := dictionary[ourMoveCode]
	enemyMove := dictionary[enemyMoveCode]

	return enemyMove, ourMove
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
