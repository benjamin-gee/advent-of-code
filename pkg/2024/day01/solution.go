package day01

import (
	"math"
	"sort"
)

func SolvePart1(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	var difference float64
	for i := 0; i < len(left); i++ {
		var distance int
		distance = right[i] - left[i]
		difference += math.Abs(float64(distance))
	}
	return int(difference)
}

func SolvePart2(left []int, right []int) int {
	counts := make(map[int]int)
	for _, num := range right {
		counts[num]++
	}

	var similarityScore int
	for _, num := range left {
		matches := counts[num] // Get the count of num from the right array
		similarityScore += num * matches
	}

	return similarityScore
}
