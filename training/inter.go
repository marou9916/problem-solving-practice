package training

func Inter(firstString string, secondString string) string {
	var result string //la string contenant les lettres communes aux deux strings dans l'ordre de la première sans double

	for _, letterinfirststring := range firstString {
		//Si la lettre est dans la seconde string et n'est pas encore rangée dans le résultat
		if IsInTheString(secondString, letterinfirststring) && !IsInTheString(result, letterinfirststring) {
			result += string(letterinfirststring) //concaténer dans la variable à retourner
		}

	}

	//Retourner les lettres communes
	return result
}

// Vérifier si une lettre est dans une string
func IsInTheString(characters string, r rune) bool {
	for _, letter := range characters {
		if letter == r {
			return true
		}
	}
	return false
}
