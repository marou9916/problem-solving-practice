package training

import (
	"os"
)

/*Ce programme transforme une chaîne de caractères passée en argument en sa version Pig Latin.

Les règles utilisées par Pig Latin sont les suivantes :

    Si un mot commence par une voyelle, ajoutez simplement "ay" à la fin.
    S'il commence par une consonne, alors nous prenons toutes les consonnes avant la première voyelle et les plaçons à la fin du mot, puis ajoutons "ay" à la fin.
    Seules les voyelles latines seront considérées comme des voyelles (aeiou).

Si le mot ne contient pas de voyelles, le programme doit afficher "Pas de voyelles".
**/

func PigLatin(word string) {
	var firstConsonants string
	var wordAfterFirstVowel string
	var indexFirstVowel int

	if IsVowel(rune(word[0])) { //Si la première lettre est une voyelle
		os.Stdout.WriteString(word + "ay" + "\n") //Rajouter "ay" à la fin du mot suivi d'un saut de ligne
	} else { //Si la première lettre est une consonne

		//Vérifier si le mot contient une voyelle
		for position, letter := range word {
			if IsVowel(letter) { //Si le mot contient une voyelle
				//Récupérer les consonnes avant la première voyelle
				firstConsonants = GetConsonantsBeforeFirstVowel(word)

				//Récupérer l'indice de la première voyelle
				indexFirstVowel = GetFisrtVowelIndex(word)

				//Récupérer le mot à partir de la première voyelle
				for j := indexFirstVowel; j <= len(word)-1; j++ {
					wordAfterFirstVowel += string(word[j])
				}

				//Afficher
				os.Stdout.WriteString(wordAfterFirstVowel + firstConsonants + "ay" + "\n")
				return
			} else if position == len(word)-1 {
				os.Stdout.WriteString("No vowels\n")
				return
			}
		}

	}
}

// Pour vérifier si le mot contient une voyelle
func IsVowel(letter rune) bool {
	return (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U')
}

// Pour récupérer les consonnes avant la première voyelle
func GetConsonantsBeforeFirstVowel(word string) string {
	var consonants string

	for _, letter := range word {
		if !IsVowel(letter) {
			consonants += string(letter)
		} else {
			break
		}
	}

	return consonants
}

// Pour récupérer l'indice de la première voyelle
func GetFisrtVowelIndex(word string) int {
	var index int

	for i := 1; i <= len(word)-1; i++ {
		if IsVowel(rune(word[i])) {
			index = i
			break
		}
	}
	return index
}
