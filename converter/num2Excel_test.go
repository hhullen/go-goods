package converter

import (
	"testing"
)

func TestNum2Excel(t *testing.T) {
	t.Parallel()

	tests := map[string]uint64{
		"A":              1,
		"Z":              26,
		"AA":             27,
		"AB":             28,
		"AZ":             52,
		"BA":             53,
		"ZY":             701,
		"ZZ":             702,
		"AAA":            703,
		"AEL":            818,
		"AWU":            1295,
		"BUE":            1903,
		"FXSHRXW":        2147483647,
		"CRPXNLSKVLJFHG": 9223372036854775807,
	}

	for expected, input := range tests {
		result := num2Excel(input)
		if result != expected {
			t.Fatalf("expected %s, got %s", expected, result)
		}
	}
}
