package piscine

func StrLen(s string) int {
	s1 := []rune(s)
	var runesCounter int
	for i := 0; i <= len(s1)-1; i++ {
		runesCounter++
	}
	return runesCounter
}

func IntToRune(the_number int) []rune {
	div := 10
	var reste int
	var digits_of_the_number []rune
	for the_number != 0 {
		reste = the_number % div
		digits_of_the_number = append(digits_of_the_number, rune(reste+'0'))
		the_number = the_number / div
	}
	return digits_of_the_number

}
