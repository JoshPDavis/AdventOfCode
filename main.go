package main

import (
	"flag"
	"github.com/joshpdavis/AOC/helpers/go"
)

func main() {

	file := flag.String("file", "input.txt", "file to run against")

	items := helpers.ReadFile(*file, func(s string) any {
		return s
	})

	print(items)
}
