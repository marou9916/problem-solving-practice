package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	if n < 0 || n > len(s)-1 {
		return 0
	} else {
		for index := range s {
			if index == n {
				index = n-1
			}
			break
		}
		return s1[index]
	}
}
