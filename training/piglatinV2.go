package training

import "os"

func PigLatinV2(word string) {
	var indexFirstVowel int
	var wordBeforeFirstVowel string
	var consonants string

	if IsVowel(rune(word[0])) {
		os.Stdout.WriteString(word + "ay" + "\n")
		return
	} else {
		for position, letter := range word {
			if IsVowel(letter) {
				indexFirstVowel = GetFirstVowelIndex(word)

				wordBeforeFirstVowel = GetWordBeforeFirstVowel(word, indexFirstVowel)

				consonants = GetConsonants(word, indexFirstVowel)

				os.Stdout.WriteString(wordBeforeFirstVowel + consonants + "ay" + "\n")
				return

			} else if position == len(word)-1 {
				os.Stdout.WriteString("No Vowels\n")
				return
			}

		}

	}
}

func GetFirstVowelIndex(word string) int {
	var index int

	for i, letter := range word {
		if IsVowel(letter) {
			index = i
			break
		}
	}
	return index
}

func GetWordBeforeFirstVowel(word string, index int) string {
	var result string

	for i := index; i <= len(word)-1; i++ {
		result += string(word[i])
	}
	return result
}

func GetConsonants(word string, index int) string {
	var consonants string

	for i := 0; i < index; i++ {
		consonants += string(word[i])
	}

	return consonants
}
