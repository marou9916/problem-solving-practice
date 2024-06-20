package training

import "github.com/01-edu/z01"

func RepeatAlphaV2(characters string) {
	for _, character := range characters {
		index := GetIndexInAlphabet(character)
		for i := 0; i <= index; i++ {
			z01.PrintRune(character)
		}
	}
	z01.PrintRune(10)
}

func GetIndexInAlphabet(character rune) int {
	for index, letter := range "abcdefghijklmnopqrstuvwxyz" {
		if character == letter {
			return index
		}
	}

	for index, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		if character == letter {
			return index
		}
	}
	return 0
}
