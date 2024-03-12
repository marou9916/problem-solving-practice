package training

func Union(string1 string, string2 string) string {
	var result string

	//Parcourir la string1
	for _, letter1 := range string1 {
		if IsInTheString(string1, letter1) && !IsInTheString(result, letter1) {
			result += string(letter1)
		}
	}

	//Parcourir la string2
	for _, letter2 := range string2 {
		if IsInTheString(string2, letter2) && !IsInTheString(result, letter2) {
			result += string(letter2)
		}
	}

	//Retourner le résultat
	return result
}
