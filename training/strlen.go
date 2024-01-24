package training

func StrLen(s string) int {
	s1 := []rune(s)
	var runesCounter int
	for i := 0; i <= len(s1)-1; i++ {
		runesCounter++
	}
	return runesCounter
}

func IntToRune(number int) []rune {
	div := 10
	var reste int
	var digits_of_the_number []rune
	for number > 0 {
		reste = number % div
		digits_of_the_number = append(digits_of_the_number, rune(reste+'0'))
		number = number / div
	}
	return digits_of_the_number
}
