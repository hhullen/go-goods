package calculation

import (
	"testing"
)

func TestGeneratePascalTriangle(t *testing.T) {
	t.Parallel()

	tests := map[int][][]int{
		0: {},
		1: {{1}},
		2: {{1}, {1, 1}},
		3: {{1}, {1, 1}, {1, 2, 1}},
		4: {{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		5: {{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		6: {{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}},
	}

	for input, expected := range tests {
		result := GeneratePascalTriangle(input)
		if len(result) != len(expected) || !cmpPascalTriangle(result, expected) {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}

func TestGenerateRow(t *testing.T) {
	t.Parallel()

	tests := map[int][]int{
		0: {1},
		1: {1, 1},
		2: {1, 2, 1},
		3: {1, 3, 3, 1},
		4: {1, 4, 6, 4, 1},
		5: {1, 5, 10, 10, 5, 1},
	}

	for input, expected := range tests {
		result := GenerateRow(input)
		if !cmpSlices(result, expected) {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}

func cmpPascalTriangle(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for row := range a {
		if !cmpSlices(a[row], b[row]) {
			return false
		}
	}
	return true
}

func cmpSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
