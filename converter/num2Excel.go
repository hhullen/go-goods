package converter

import (
	"strings"
)

func num2Excel(columnNumber uint64) string {
	var sb strings.Builder
	sb.Grow(4)
	var recF func(uint64)
	recF = func(n uint64) {
		mod := n % 26
		if n > 26 {
			if mod == 0 {
				recF(n/26 - 1)
			} else {
				recF(n / 26)
			}
		}

		if mod == 0 {
			sb.WriteByte('Z')
			return
		}

		sb.WriteByte(byte(mod + 64))
	}

	recF(columnNumber)

	return sb.String()
}
