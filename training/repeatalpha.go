package training

func RepeatAlpha(s string) string {
	var result string //la string à retourner
	var index int
	//Convertir la string en runes
	s1 := []rune(s)

	//Parcourir la string
	for i := 0; i <= len(s1)-1; i++ {
		if (s1[i] >= 'A' && s1[i] <= 'Z') || (s1[i] >= 'a' && s1[i] <= 'z') { //Si on tombe sur une lettre
			//Récupérer l'index de la lettre dans l'alphabet
			index = AlphaIndex(s1[i])

			//Concaténer la lettre dans la variable autant de fois que son index
			for j := 1; j <= index; j++ {
				result += string(s1[i]) //Répéter la lettre dans la string autant de fois que son rang dans l'alphabet
			}
		} else {
			result += string(s1[i])
		}
	}
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
