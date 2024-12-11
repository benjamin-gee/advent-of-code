package main

import (
	"bufio"
	"fmt"
	"github.com/benjamin-gee/advent-of-code/pkg/2024/day02"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("pkg/2024/day02/input.txt")
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

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		numbers := make([]int, 0)

		for _, field := range fields {
			num, err := strconv.Atoi(field)

			if err != nil {
				fmt.Printf("error parsing field: %v\n", err)
				continue
			}

			numbers = append(numbers, num)

		}

		reports = append(reports, numbers)
	}

	var safeReports = day02.SolvePart1(reports)

	fmt.Printf("Day 2 Solve Part 1: %d\n", safeReports)

	var moderatelySafeReports = day02.SolvePart2(reports)

	fmt.Printf("Day 2 Solve Part 1: %d\n", moderatelySafeReports)
}
