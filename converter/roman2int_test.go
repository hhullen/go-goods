package converter

import "testing"

func TestRomanToInt(t *testing.T) {
	t.Parallel()

	tests := map[string]int{
		"I":       1,
		"II":      2,
		"III":     3,
		"IV":      4,
		"V":       5,
		"VI":      6,
		"VII":     7,
		"VIII":    8,
		"IX":      9,
		"X":       10,
		"LVIII":   58,
		"MCMXCIV": 1994,
		"MMXXVI":  2026,
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDLXV":                                     65565,
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMDCLIV":              87654,
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMCMXCIX": 99999,
	}

	for input, expected := range tests {
		actual := RomanToInt(input)
		if actual != expected {
			t.Errorf("RomanToInt(%q) = %d; want %d", input, actual, expected)
		}
	}
}
