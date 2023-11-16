package piscine

func StrLen(s string) int {
	s1 := []rune(s)
	runeCounter := 0
	for i := 0; i <= len(s1)-1; i++ {
		runeCounter++
	}
	return runeCounter
}
