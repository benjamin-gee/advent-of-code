package day02

import "math"

func SolvePart1(reports [][]int) int {
	var safeReports = 0

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func isSafe(report []int) bool {
	var increasing = false

	if report[0] < report[1] {
		increasing = true
	}

	for i := 1; i < len(report); i++ {
		current := report[i]
		previous := report[i-1]

		var difference = current - previous

		if increasing && difference < 1 {
			return false
		}

		if !increasing && difference > -1 {
			return false
		}

		if math.Abs(float64(difference)) > 3 {
			return false
		}
	}

	return true
}

func SolvePart2(reports [][]int) int {
	var safeReports = 0

	for _, report := range reports {
		if isModeratelySafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func isModeratelySafe(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		reportWithoutI := append([]int{}, report[:i]...)
		reportWithoutI = append(reportWithoutI, report[i+1:]...)
		if isSafe(reportWithoutI) {
			return true
		}
	}

	return false
}
