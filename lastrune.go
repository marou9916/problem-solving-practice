package piscine

func LastRune(Characters string) rune {
	convertedCharacters := []rune(Characters)
	return convertedCharacters[len(convertedCharacters)-1]
}




