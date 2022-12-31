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

	packetIndex, messageIndex := findSignalStartIndex(inputFile)

	fmt.Println("Packet sequence start index is", packetIndex+1)
	fmt.Println("Message sequence start index is", messageIndex+1)

}

var PACKET_SEQUENCE_LENGTH = 4
var MESSAGE_SEQUENCE_LENGTH = 14

func findSignalStartIndex(inputFile *os.File) (int, int) {
	packetSequenceIndex := 0
	messageSequenceIndex := 0
	packetSequence := ""
	messageSequence := ""
	forCharOfFile(inputFile, func(i int, c string) error {
		if len(packetSequence) == PACKET_SEQUENCE_LENGTH {
			packetSequence = packetSequence[1:]
		}
		if len(messageSequence) == MESSAGE_SEQUENCE_LENGTH {
			messageSequence = messageSequence[1:]
		}

		packetSequence += c
		messageSequence += c

		if packetSequenceIndex == 0 && checkSequence(packetSequence, PACKET_SEQUENCE_LENGTH) {
			packetSequenceIndex = i
		}

		if messageSequenceIndex == 0 && checkSequence(messageSequence, MESSAGE_SEQUENCE_LENGTH) {
			messageSequenceIndex = i
			return io.EOF
		}

		return nil

	})
	return packetSequenceIndex, messageSequenceIndex
}

func checkSequence(s string, length int) bool {
	if len(s) < length {
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
