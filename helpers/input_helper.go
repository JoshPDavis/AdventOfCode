package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile[K any](path string, stringToObj func(s string) K) []K {
	items := []K{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, stringToObj(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return items
}
