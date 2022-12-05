package main

import (
	"flag"
	"log"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

func main() {

	file := flag.String("file", "input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) int {
		a := s[:(len(s) / 2)]
		b := s[len(s)/2:]
		return SumIntersection(Intersection(a, b))
	})

	sum := 0
	for _, v := range items {
		sum += v
	}

	log.Printf("Total:%v", sum)
}

func Intersection(a string, b string) map[byte]bool {
	intersect := map[byte]bool{}
	seen := map[byte]bool{}
	for index, _ := range a {
		seen[a[index]] = true
	}

	for index, _ := range b {
		_, keyExist := seen[b[index]]
		if keyExist {
			intersect[b[index]] = true
		}
	}

	log.Print(intersect)
	return intersect
}

func SumIntersection(i map[byte]bool) int {
	sum := 0
	for key, _ := range i {
		sum += GetValue(key)
	}
	return sum
}

func GetValue(a byte) int {
	ascii := int(a)
	if ascii >= 97 {
		return ascii - 96
	} else {
		return ascii - 38
	}
}
