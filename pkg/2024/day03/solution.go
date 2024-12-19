package day03

import (
	"fmt"
	"regexp"
)

func SolvePart1(corruptedMemory string) int {
	var sum = 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := re.FindAllString(corruptedMemory, -1)

	for _, match := range matches {
		var x, y int
		_, err := fmt.Sscanf(match, "mul(%d,%d)", &x, &y)

		if err != nil {
			fmt.Printf("error parsing match: %v\n", err)
		}
		sum = sum + (x * y)
	}

	return sum
}

func SolvePart2(corruptedMemory string) int {
	var sum = 0

	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\))`)

	matches := re.FindAllString(corruptedMemory, -1)

	do := true

	for _, match := range matches {
		switch match {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				var x, y int
				_, err := fmt.Sscanf(match, "mul(%d,%d)", &x, &y)

				if err != nil {
					fmt.Printf("error parsing match: %v\n", err)
				}
				sum = sum + (x * y)
			}
		}
	}

	return sum
}
