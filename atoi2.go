package piscine

func Atoi2(s string) int {
	result := 0
	decimal := 1
	for i := len(s) - 1; i >= 0; i++ {
		if s[i] == '-' {
			result *= -1
		} else {
			result += int(s[i]+48) * decimal
			decimal *= 10
		}
	}
	return result
}
