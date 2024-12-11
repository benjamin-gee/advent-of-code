package day02

import "testing"

func TestSolvePart1(t *testing.T) {
	tests := []struct {
		name     string
		reports  [][]int
		expected int
	}{
		{
			name: "Test 1",
			reports: [][]int{
				{1, 2, 3},  // Safe Case
				{10, 9, 7}, // Safe Case
				{2, 4, 7},  // Safe Case
				{1, 1, 3},
				{10, 2, 3},
				{1, 2, 2},
				{1, 2, 6},
				{1, 2, -3},
				{3, 3, 3}},
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SolvePart1(test.reports)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name     string
		reports  [][]int
		expected int
	}{
		{
			name: "Test 1",
			reports: [][]int{
				{7, 6, 4, 2, 1},                  // Safe
				{1, 2, 7, 8, 9},                  // Unsafe
				{9, 7, 6, 2, 1},                  // Unsafe
				{1, 3, 2, 4, 5},                  // Safe
				{8, 6, 4, 4, 1},                  // Safe
				{1, 3, 6, 7, 9},                  // Safe
				{14, 17, 20, 21, 24, 26, 27, 24}, // Safe
				{39, 41, 43, 45, 46, 46},         // Safe
				{35, 38, 39, 41, 44, 47, 50, 54}, // Safe
				{68, 69, 71, 74, 75, 78, 80, 87}, // Safe
				{80, 82, 81, 82, 83, 85, 88},     // Safe
				{48, 51, 54, 55, 58, 57, 55},     // Unsafe
				{41, 44, 47, 50, 47, 47},         // Unsafe
				{66, 68, 71, 70, 73, 77},         // Unsafe
				{29, 32, 29, 30, 35},             // Unsafe
				{77, 78, 79, 81, 83, 83, 86, 88}, // Safe
				{10, 13, 14, 16, 19, 19, 20, 17}, // Unsafe
				{61, 64, 65, 67, 67, 67},         // Unsafe
				{29, 30, 31, 32, 35, 35, 39},     // Unsafe
				{24, 25, 25, 28, 31, 38},         // Unsafe
				{61, 64, 65, 66, 70, 73, 76},     // Unsafe
				{35, 37, 39, 43, 40},             // Safe
				{41, 42, 43, 47, 48, 49, 49},     // Unsafe
				{3, 5, 9, 11, 15}},               // Unsafe
			expected: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SolvePart2(test.reports)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
