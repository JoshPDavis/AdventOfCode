package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

type Range struct {
	str        string
	lowerBound int
	upperBound int
}

func NewRange(s string) *Range {
	nums := strings.Split(s, "-")

	l, _ := strconv.Atoi(nums[0])
	u, _ := strconv.Atoi(nums[1])

	return &Range{
		lowerBound: l,
		upperBound: u,
	}
}

func main() {

	file := flag.String("file", "../input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) [2]Range {
		ranges := strings.Split(s, ",")
		return [2]Range{*NewRange(ranges[0]), *NewRange(ranges[1])}
	})

	total := 0
	for _, sections := range items {
		if overlap(sections[0], sections[1]) {
			total += 1
		}
	}

	log.Printf("Total: %v", total)
}

func overlap(s1, s2 Range) bool {
	a := s1.lowerBound
	b := s1.upperBound

	c := s2.lowerBound
	d := s2.upperBound
	if (c < a && d < a && d < b && c < b) || (a < c && b < c && a < d && b < d) {
		return false
	}
	return true
}
