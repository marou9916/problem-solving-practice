package piscine

func AlphaCount(s string) int {
	count := 1
	s1 := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if s1[i] >= 'a' && s1[i] <= 'z' || s1[i] >= 'A' && s1[i] <= 'Z' || s1[i] < '0' && s1[i] > '9' {
			count = count + 1
		}
	}
	return count
}
