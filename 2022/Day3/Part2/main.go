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

	sum := 0
	for i := 0; i <= len(items)-3; i += 3 {
		a := items[i]
		b := items[i+1]
		c := items[i+2]

		sum += SumIntersection(Intersection(a, b, c))
	}

	log.Printf("Total:%v", sum)
}

func Intersection(a string, b string, c string) map[byte]bool {
	intersect := map[byte]bool{}
	seen := map[byte]int{}
	for index := range a {
		seen[a[index]] = 1
	}

	for index := range b {
		_, keyExist := seen[b[index]]
		if keyExist {
			seen[b[index]] += 1
		}
	}

	for index := range c {
		v, keyExist := seen[c[index]]
		if keyExist && v >= 2 {
			intersect[c[index]] = true
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
