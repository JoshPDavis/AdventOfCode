package main

import (
	"flag"

	helpers "github.com/joshpdavis/AOC/helpers/go"
)

func main() {

	file := flag.String("file", "input.txt", "file to run against")
	flag.Parse()

	items := helpers.ReadFile(*file, func(s string) any {
		return s
	})

	print(items)
}
