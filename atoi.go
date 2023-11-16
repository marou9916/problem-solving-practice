package piscine

func Atoi(s string) int {
	s1 := []rune(s)
	signCounter := 0

	for _, char := range s1 {
		if char == ' ' {
			return 0
		}
	}
	for _, char := range s1 {
		if 'a' < char && char < 'z' || 'A' < char && char < 'Z' {
			return 0
		}
	}
	for _, char := range s1 {
		if char == '+' || char == '-' {
			signCounter += 1
			if signCounter == 2 {
				return 0
			}
		}
	}
	if s1[0] == '+' || s1[0] == '-' {
		for i := 1; i <= len(s1)-1; i++ {
			if '0' <= s1[i] && s1[i] <= '9' {
				return int(s1[i] - 48)
			}
		}
	}
	return 0
}
