//package main

import "github.com/01-edu/z01"

func TestRot14(s string) string {
	s1 := []rune(s)
	var result string
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'a' && s1[i] < 'n' || s1[i] >= 'A' && s1[i] < 'N' {
			s1[i] += 14
		} else if s1[i] == 'n' {
			s1[i] = 'a'
		} else if s1[i] == 'N' {
			s1[i] = 'A'
		} else if s1[i] >= 'o' && s1[i] <= 'z' || s1[i] >= 'O' && s1[i] < 'Z' {
			s1[i] += 12
		}
	}
	result += string(s1)
	return string(result)
}

func main() {
	result := TestRot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
