package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

type Stack []rune

func (s Stack) Push(letter rune) Stack {
	return append(s, letter)
}

func (s Stack) Pop() (Stack, rune) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func ReadInstruction(s string) (ops, from, to int) {
	instructions := strings.Fields(s)

	ops, _ = strconv.Atoi(instructions[1])
	from, _ = strconv.Atoi(instructions[3])
	to, _ = strconv.Atoi(instructions[5])
	return
}

func main() {

	file := flag.String("file", "../input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) string {
		return s
	})

	crates := map[int]Stack{
		1: []rune{'D', 'H', 'N', 'Q', 'T', 'W', 'V', 'B'},
		2: []rune{'D', 'W', 'B'},
		3: []rune{'T', 'S', 'Q', 'W', 'J', 'C'},
		4: []rune{'F', 'J', 'R', 'N', 'Z', 'T', 'P'},
		5: []rune{'G', 'P', 'V', 'J', 'M', 'S', 'T'},
		6: []rune{'B', 'W', 'F', 'T', 'N'},
		7: []rune{'B', 'L', 'D', 'Q', 'F', 'H', 'V', 'N'},
		8: []rune{'H', 'P', 'F', 'R'},
		9: []rune{'Z', 'S', 'M', 'B', 'L', 'N', 'P', 'H'},
	}

	// crates := map[int]Stack{
	// 	1: []rune{'Z', 'N'},
	// 	2: []rune{'M', 'C', 'D'},
	// 	3: []rune{'P'},
	// }
	for i := 10; i < len(items); i++ {
		ops, from, to := ReadInstruction(items[i])

		// log.Printf("ops:%v from:%v to:%v", ops, from, to)
		for j := 0; j < ops; j++ {
			stack, v := crates[from].Pop()
			crates[from] = stack
			s2 := crates[to].Push(v)
			crates[to] = s2
		}
	}

	list := []rune{}

	for i := 1; i <= 9; i++ {
		list = append(list, crates[i][len(crates[i])-1])
	}

	log.Printf("Final Crate Position:%c", crates)
	log.Printf("Final Crate Position:%c", list)
}
