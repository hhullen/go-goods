package converter

var dict = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	res := dict[s[len(s)-1]]

	for i := len(s) - 1; i > 0; i-- {
		if dict[s[i-1]] < dict[s[i]] {
			res -= dict[s[i-1]]
		} else {
			res += dict[s[i-1]]
		}
	}

	return res
}
