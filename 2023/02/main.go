package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hpdobrica/advent-of-code/util"
)

func main() {
	start := time.Now()

	file, err := os.Open("./input.txt")

	util.PanicIfErr(err)

	var wg sync.WaitGroup
	limiter := make(chan struct{}, 10)
	sumChan := make(chan int)

	sumOfValidIds := 0

	util.ForLineOfFile(file, func(line string) {
		limiter <- struct{}{}
		wg.Add(1)
		go processLine(line, &wg, sumChan)
		<-limiter
	})

	go func() {
		wg.Wait()
		close(sumChan)
	}()

	for id := range sumChan {
		sumOfValidIds += id
	}

	fmt.Println("Sum of valid game ids is: ", sumOfValidIds)
	elapsed := time.Since(start)
	fmt.Println("Elapsed: ", elapsed)

}

type GameData struct {
	ID     int
	Rounds []RoundData
}

type RoundData struct {
	Red   int
	Green int
	Blue  int
}

func processLine(line string, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- getIdOfValidGameOrZero(line)

}

func getIdOfValidGameOrZero(line string) int {

	gameData := lineToGameData(line)

	if isGameValid(gameData) {
		return gameData.ID
	}
	return 0
}

func isGameValid(game GameData) bool {
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	MAX_RED := 12
	MAX_GREEN := 13
	MAX_BLUE := 14

	for _, round := range game.Rounds {
		if round.Red > MAX_RED || round.Green > MAX_GREEN || round.Blue > MAX_BLUE {
			return false
		}
	}
	return true
}

func lineToGameData(line string) GameData {
	splitByColon := strings.Split(line, ":")

	idString := splitByColon[0][len("Game: ")-1:]

	id, err := strconv.Atoi(idString)

	if err != nil {
		panic("game id is not a number!")
	}

	roundDataStrings := strings.Split(strings.TrimSpace(splitByColon[1]), ";")
	rounds := make([]RoundData, len(roundDataStrings))

	for roundIndex, singleRoundData := range roundDataStrings {
		roundParts := strings.Split(singleRoundData, ",")
		newRound := RoundData{}
		for _, part := range roundParts {
			splitPart := strings.Split(strings.TrimSpace(part), " ")
			partName := splitPart[1]
			partValue, err := strconv.Atoi(splitPart[0])
			if err != nil {
				panic("part value is not a number!")
			}

			if partName == "red" {
				newRound.Red = partValue

			}
			if partName == "green" {
				newRound.Green = partValue
			}
			if partName == "blue" {
				newRound.Blue = partValue
			}
		}
		rounds[roundIndex] = newRound

	}

	return GameData{
		ID:     id,
		Rounds: rounds,
	}

}
