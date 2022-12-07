package main

import (
	"flag"
	"log"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

func main() {

	file := flag.String("file", "input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) string {
		return s
	})

	var position int
	uniqueLetters := map[rune]bool{}
	for index, v := range items[0] {
		_, keyExists := uniqueLetters[v]
		if !keyExists {
			uniqueLetters[v] = true
			if len(uniqueLetters) >= 4 {
				log.Printf("Char is %c", v)
				position = index
				break
			}
		} else {
			uniqueLetters = map[rune]bool{}
		}
	}

	log.Printf("5th Char is: %v", position)
}
