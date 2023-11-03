package piscine

func AlphaCount(s string) int {
	count := 0
	
	for i := 0; i <= len(s) - 1; i++ {
		if s[i] > 65 && s[i] < 90 {
			count = count + 1
			
		}
	}
	return count
}
