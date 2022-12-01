package main

import (
	"github.com/joshpdavis/AOC/helpers"
)

func main() {

	items := helpers.ReadFile("input.txt", func(s string) any {
		return s
	})

	print(items)
}
