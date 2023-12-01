package main

import (
	"log"
	"os"
	"strings"
)

const filePath = "input.txt"

func main() {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("could not found %v", filePath)
	}

	lines := strings.Split(string(bytes), "\n")

	sum := 0

	for _, line := range lines {
		values := []int{}
		for _, ch := range line {
			if 48 <= ch && ch <= 57 {
				values = append(values, int(ch-48))
			}
		}
		if len(values) > 0 {
			sum += 10 * values[0]
			sum += values[len(values)-1]
		}
	}

	log.Print(sum)
}
