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

type Outcome struct {
	name  string
	score int
}

var ROCK = Move{"ROCK", 1, nil}
var PAPER = Move{"PAPER", 2, nil}
var SCISSORS = Move{"SCISSORS", 3, nil}

var LOSE = Outcome{"LOSE", 0}
var DRAW = Outcome{"DRAW", 3}
var WIN = Outcome{"WIN", 6}

func main() {
	ROCK.winsAgainst = &SCISSORS
	PAPER.winsAgainst = &ROCK
	SCISSORS.winsAgainst = &PAPER

	file, err := os.Open("./input.txt")
	check(err)

	totalFakeScore, totalRealScore := calculateScores(file)

	fmt.Println("We first thought total score is", totalFakeScore, ", but the real score is actually", totalRealScore)

}

func calculateScores(file *os.File) (int, int) {

	totalFakeScore := 0
	totalRealScore := 0

	forLineOfFile(file, func(line string) {
		totalFakeScore += calculateFakeScoreForRound(line[0:1], line[2:3])
		totalRealScore += calculateRealScoreForRound(line[0:1], line[2:3])
	})

	return totalFakeScore, totalRealScore
}

// part two

func calculateRealScoreForRound(enemyMoveCode string, ourMoveCode string) int {

	roundScore := 0

	var ourMove Move
	enemyMove, intendedOutcome := realMoveDecode(enemyMoveCode, ourMoveCode)

	if intendedOutcome == DRAW {
		ourMove = enemyMove
	} else if intendedOutcome == LOSE {
		ourMove = *enemyMove.winsAgainst
	} else {
		ourMove = *enemyMove.winsAgainst.winsAgainst
	}

	roundScore += ourMove.score
	roundScore += intendedOutcome.score

	return roundScore

}

func realMoveDecode(enemyMoveCode string, intendedOutcomeCode string) (Move, Outcome) {
	var enemyMoves = map[string]Move{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
	}

	var intendedOutcomes = map[string]Outcome{
		"X": LOSE,
		"Y": DRAW,
		"Z": WIN,
	}

	enemyMove := enemyMoves[enemyMoveCode]
	intendedOutcome := intendedOutcomes[intendedOutcomeCode]

	return enemyMove, intendedOutcome
}

// part one

func calculateFakeScoreForRound(enemyMoveCode string, ourMoveCode string) int {

	roundScore := 0

	enemyMove, ourMove := fakeMoveDecode(enemyMoveCode, ourMoveCode)

	roundScore += ourMove.score

	if ourMove == enemyMove {
		roundScore += DRAW.score
	} else if *ourMove.winsAgainst == enemyMove {
		roundScore += WIN.score
	} else {
		roundScore += LOSE.score
	}

	return roundScore

}

func fakeMoveDecode(enemyMoveCode string, ourMoveCode string) (Move, Move) {
	var dictionary = map[string]Move{
		"A": ROCK,
		"X": ROCK,
		"B": PAPER,
		"Y": PAPER,
		"C": SCISSORS,
		"Z": SCISSORS,
	}

	enemyMove := dictionary[enemyMoveCode]
	ourMove := dictionary[ourMoveCode]

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
