package main

import (
	"flag"
	"log"
	"strings"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

type ElfChoice string

const (
	ElfRock     ElfChoice = "A"
	ElfPaper    ElfChoice = "B"
	ElfScissors ElfChoice = "C"
)

type MyChoice string

const (
	MyRock     MyChoice = "X"
	MyPaper    MyChoice = "Y"
	MyScissors MyChoice = "Z"
)

var MyChoiceToElfChoice = map[MyChoice]ElfChoice{
	MyRock:     ElfRock,
	MyPaper:    ElfPaper,
	MyScissors: ElfScissors,
}

func GetChoicePointValue(m ElfChoice) int {
	switch {
	case m == ElfRock:
		return 1
	case m == ElfPaper:
		return 2
	case m == ElfScissors:
		return 3
	default:
		return 0
	}
}

func main() {

	file := flag.String("file", "input.txt", "file to run against")

	items := helpers.ReadFile(*file, func(s string) int {
		strArr := strings.Split(s, " ")

		return GetRoundPoints(ElfChoice(strArr[0]), MyChoiceToElfChoice[MyChoice(strArr[1])])
	})

	sum := 0
	for _, round := range items {
		sum += round
	}

	log.Printf("Total Points: %v", sum)

}

func GetRoundPoints(elf, me ElfChoice) int {
	if elf == ElfRock && me == ElfPaper || elf == ElfPaper && me == ElfScissors || elf == ElfScissors && me == ElfRock {
		return 6 + GetChoicePointValue(me)
	} else if elf == me {
		return 3 + GetChoicePointValue(me)
	} else {
		return 0 + GetChoicePointValue(me)
	}
}
