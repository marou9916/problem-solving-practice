package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	} else {
		return s1[n-1]
	}
}