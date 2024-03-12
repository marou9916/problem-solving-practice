package training

import "github.com/01-edu/z01"

func AlphaMirror(characters string) {
	//Déclaration de variables
	runes_of_the_characters := []rune(characters) //Conversion des caractères en runes
	var oppositeLetter rune                       // Variable de stockage de l'opposé de la lettre

	//Parours de la chaîne de caractères
	for i := range runes_of_the_characters {
		//Si le caractère est une lettre
		if IsLetter(runes_of_the_characters[i]) {
			//Remplacer la lettre par son opposé
			oppositeLetter = ReplaceLetterByItsOpposite(runes_of_the_characters[i])
			//Afficher son opposé
			z01.PrintRune(oppositeLetter)
		} else {
			//Sinon afficher le caractère
			z01.PrintRune(runes_of_the_characters[i])
		}
	}
	//Saut de ligne
	z01.PrintRune(10)

}

// Pour vérifier si le caractère est une lettre
func IsLetter(character rune) bool {
	return character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z'
}

// Pour remplacer la lettre par son opposé
func ReplaceLetterByItsOpposite(letter rune) rune {
	var opposite_of_the_letter rune

	if letter >= 'A' && letter <= 'Z' {
		for i := 'A'; i <= 'Z'; i++ {
			if letter == i {
				opposite_of_the_letter = 'A' + 'Z' - letter
				break
			}
		}
	}

	if letter >= 'a' && letter <= 'z' {
		for i := 'a'; i <= 'z'; i++ {
			if letter == i {
				opposite_of_the_letter = 'a' + 'z' - letter
				break
			}
		}

	}

	return opposite_of_the_letter

}
