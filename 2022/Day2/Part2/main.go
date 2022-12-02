package main

import (
	"flag"
	"log"
	"strings"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

type Choice string

const (
	ElfRock     Choice = "A"
	ElfPaper    Choice = "B"
	ElfScissors Choice = "C"
)

var winningPair = map[Choice]Choice{
	ElfPaper:    ElfScissors,
	ElfRock:     ElfPaper,
	ElfScissors: ElfRock,
}

var losingPair = map[Choice]Choice{
	ElfPaper:    ElfRock,
	ElfRock:     ElfScissors,
	ElfScissors: ElfPaper,
}

type MyStrategy string

const (
	Lose MyStrategy = "X"
	Draw MyStrategy = "Y"
	Win  MyStrategy = "Z"
)

func GetRoundPoints(m MyStrategy, elfChoice Choice) int {
	switch {
	case m == Lose:
		return GetChoicePointValue(losingPair[elfChoice]) + 0
	case m == Draw:
		return GetChoicePointValue(elfChoice) + 3
	case m == Win:
		return GetChoicePointValue(winningPair[elfChoice]) + 6
	default:
		panic("Not A choice")
	}
}

func GetChoicePointValue(m Choice) int {
	switch {
	case m == ElfRock:
		return 1
	case m == ElfPaper:
		return 2
	case m == ElfScissors:
		return 3
	default:
		panic("Not A Choice")
	}
}

func main() {

	file := flag.String("file", "input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) int {
		strArr := strings.Split(s, " ")

		return GetRoundPoints(MyStrategy(strArr[1]), Choice(strArr[0]))
	})

	sum := 0
	for _, round := range items {
		sum += round
	}

	log.Printf("Total Points: %v", sum)

}
