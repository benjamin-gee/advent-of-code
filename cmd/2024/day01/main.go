package main

import (
	"bufio"
	"github.com/benjamin-gee/advent-of-code/pkg/2024/day01"
	"os"
	"strconv"
	"strings"
)
import "fmt"

func main() {
	file, err := os.Open("pkg/2024/day01/input.txt")
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

	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		numbers := strings.Fields(line)

		leftVal, _ := strconv.ParseInt(numbers[0], 10, 64)
		rightVal, _ := strconv.ParseInt(numbers[1], 10, 64)

		left = append(left, int(leftVal))
		right = append(right, int(rightVal))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error scanning file: %v\n", err)
		return
	}

	day01ResultPart1 := day01.SolvePart1(left, right)
	fmt.Printf("Day 1 Solution: %d\n", day01ResultPart1)

	day01ResultPart2 := day01.SolvePart2(left, right)
	fmt.Printf("Day 1 Solution: %d\n", day01ResultPart2)
}
