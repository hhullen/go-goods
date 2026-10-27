package converter

import "testing"

func TestExcel2Num(t *testing.T) {
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

	for columnTitle, expected := range tests {
		result := Excel2Num(columnTitle)
		if result != expected {
			t.Errorf("Excel2Num(%s) = %d; expected %d", columnTitle, result, expected)
		}
	}
}
