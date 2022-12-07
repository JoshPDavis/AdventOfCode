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

	index := 0
	uniqueLetters := map[byte]int{}
	for len(uniqueLetters) < 14 {
		v := items[0][index]
		_, keyExists := uniqueLetters[v]
		if !keyExists {
			uniqueLetters[v] = index
			index++
		} else {
			// log.Printf("Current Index: %v", index)
			index = uniqueLetters[v] + 1
			// log.Printf("REPEAT is %c, New Index is:%v", v, index)
			uniqueLetters = map[byte]int{}
		}
	}

	log.Printf("14th Char is: %v", index)
	//TODO: Make Better
	packetering(items[0])
}

func packetering(str string) {

	var position int
	lastOccurance := map[byte]int{}

	for i := 0; i < len(str); i++ {
		_, keyExists := lastOccurance[str[i]]
		if !keyExists {
			lastOccurance[str[i]] = i
			if len(lastOccurance) >= 14 {
				position = i
				break
			}
		} else {
			lastOccurance = map[byte]int{}
			position = i
		}
	}

	log.Printf("PacketIndex: %v", position)
}
