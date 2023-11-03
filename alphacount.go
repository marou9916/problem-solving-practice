package piscine

func AlphaCount(s string) int {
	count := 0
	i := 0
	for i < len(s) {
		if s[i] > 65 && s[i] < 90 {
			count = count + 1
			i++
		}
	}
	return count
}
