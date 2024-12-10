package day01

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		left     []int
		right    []int
		expected int
	}{
		{
			name:     "Test 1",
			left:     []int{5, 3, 2},
			right:    []int{1, 1, 1},
			expected: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SolvePart1(test.left, test.right)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestSolutionPart2(t *testing.T) {
	tests := []struct {
		name     string
		left     []int
		right    []int
		expected int
	}{
		{
			name:     "Test 1",
			left:     []int{1, 2, 3},
			right:    []int{3, 3, 3},
			expected: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SolvePart2(test.left, test.right)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
