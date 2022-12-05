package main

import (
	"flag"
	"log"
	"strconv"
	"strings"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

type SectionRange struct {
	lowerBound int
	upperBound int
}

func NewSectionRange(s string) *SectionRange {
	nums := strings.Split(s, "-")

	l, _ := strconv.Atoi(nums[0])
	u, _ := strconv.Atoi(nums[1])

	return &SectionRange{
		lowerBound: l,
		upperBound: u,
	}
}
func main() {

	file := flag.String("file", "input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) [2]SectionRange {
		ranges := strings.Split(s, ",")
		return [2]SectionRange{*NewSectionRange(ranges[0]), *NewSectionRange(ranges[1])}
	})

	total := 0
	for _, sections := range items {
		if (sections[0].lowerBound <= sections[1].lowerBound && sections[0].upperBound >= sections[1].upperBound) ||
			(sections[1].lowerBound <= sections[0].lowerBound && sections[1].upperBound >= sections[0].upperBound) {
			total += 1
		}
	}

	log.Printf("Total: %v", total)
}
