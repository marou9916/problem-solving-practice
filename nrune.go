package piscine

func NRune(s string, n int) rune {
	x := 0
	s1 := []rune(s)
	if n < 0 || n > len(s)-1 {
		return 0
	} else {
		for index := range s {
			if index == n {
				x = n - 1
			}
			break
		}
		return s1[x]
	}
}

