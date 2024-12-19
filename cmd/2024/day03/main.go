package main

import (
	"bufio"
	"fmt"
	"github.com/benjamin-gee/advent-of-code/pkg/2024/day03"
	"os"
)

func main() {
	file, err := os.Open("pkg/2024/day03/input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %v\n", err)
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		line := scanner.Text()
		sum := day03.SolvePart1(line)
		sumDay2 := day03.SolvePart2(line)
		fmt.Printf("Day 3 Solve Part 1: %d\n", sum)
		fmt.Printf("Day 3 Solve Part 2: %d\n", sumDay2)
	}
}
