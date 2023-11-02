package piscine

func LastRune(s string) rune {
	s1 := []rune(s)
	return s1[len(s)-1]
}
