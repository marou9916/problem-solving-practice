package piscine

func testActivebits(n int) int {
	count := 0
	for n > 0 {
		count = count + 1
		n >>= 1
	}
	return count
}
