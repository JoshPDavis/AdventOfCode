package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/joshpdavis/AOC/helpers"
)

func main() {

	items := helpers.ReadFile("input.txt", func(s string) string {
		return s
	})

	max := 0
	runningSum := 0

	sums := []int{}

	for _, snack := range items {
		calories, err := strconv.Atoi(snack)
		if err != nil {
			if max < runningSum {
				max = runningSum
				log.Print(fmt.Sprint(runningSum) + "\n")
			}
			sums = append(sums, runningSum)
			runningSum = 0
		} else {
			runningSum = calories + runningSum
		}
	}

	sort.Ints(sums)

	a := sums[len(sums)-1]
	b := sums[len(sums)-2]
	c := sums[len(sums)-3]

	log.Print(a + b + c)
	log.Print(a)
	log.Printf("Largest Sum:%v", max)
}
