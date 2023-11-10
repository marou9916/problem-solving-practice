package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		count += 1
		n >>= 1
	}
	return int(count)
}
