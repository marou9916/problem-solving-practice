package training

func RepeatAlpha(s string) string {
	var result string //la string à retourner
	var index int     //l'indice de la lettre dans l'alphabet

	//Convertir la string en runes
	s1 := []rune(s)

	//Parcourir les runes
	for i := 0; i <= len(s1)-1; i++ {
		if (s1[i] >= 'A' && s1[i] <= 'Z') || (s1[i] >= 'a' && s1[i] <= 'z') { //Si l'élément est une lettre
			//Récupérer l'index de la lettre dans l'alphabet
			index = AlphaIndex(s1[i])

			//Concaténer la lettre dans la variable à retourner autant de fois que son index
			for j := 1; j <= index; j++ {
				result += string(s1[i]) //Répéter la lettre dans la string autant de fois que son index dans l'alphabet
			}
		} else { //Sinon (si l'élément n'est pas une lettre)
			result += string(s1[i]) //Ajouter l'élément sans le répéter
		}
	}

	//Retourner le résultat
	return result
}

func AlphaIndex(letter rune) int {
	max := 26
	min := 1
	var index int

	if letter == 'a' || letter == 'A' {
		index = min
	} else if letter == 'z' || letter == 'Z' {
		index = max
	} else if letter > 'a' && letter < 'z' {
		index = max - int('z'-letter)
	} else if letter > 'A' && letter < 'Z' {
		index = max - int('Z'-letter)
	}

	return index
}

//Version 2

// func IsLetter(character rune) bool {
// 	return (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z')
// }

// func GetPositionInAlphabet(letter rune) int {
// 	var index int

// 	if letter >= 'a' && letter <= 'z' {
// 		for i := 'a'; i <= letter; i++ {
// 			index++
// 		}
// 	}

// 	if letter >= 'A' && letter <= 'Z' {
// 		for i := 'A'; i <= letter; i++ {
// 			index++
// 		}
// 	}
// 	return index
// }

// func RepeatAlpha2(letter rune) {
// 	index := GetPositionInAlphabet(letter)
// 	for i := 1; i <= index; i++ {
// 		z01.PrintRune(letter)
// 	}
// }
