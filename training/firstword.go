package training

import "unicode"

func FirstWord(input string) string {
	var output string

	for _, character := range input {
		if IsSpace(character) {
			if output == "" {
				continue
			} else {
				return output
			}
		} else {
			output += string(character)
		}
	}

	return output
}

func IsSpace(character rune) bool {
	return unicode.IsSpace(character)
}
