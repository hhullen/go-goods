package converter

func Excel2Num(columnTitle string) uint64 {
	if len(columnTitle) < 1 {
		return 0
	}

	res := uint64(0)
	base := uint64(1)
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += (uint64(columnTitle[i]) - 64) * base
		base *= 26
	}

	return res
}
