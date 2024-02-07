package training

func Itoa(the_int int) string {
	//Déclaration des variables
	var runesOfNumber []rune
	number := the_int
	rest := 0
	result := ""

	//Gérer le cas du nombre 0
	if number == 0 {
		return "0"
	}

	//Gérer le signe du nombre négatif
	if number < 0 {
		result += "-"
		number *= -1
	}

	//Récupérer les digits du nombre en runes dans un tableau de runes
	for number > 0 {
		rest, number = number%10, number/10
		runesOfNumber = append(runesOfNumber, rune(rest+'0'))
	}

	//Concaténer les runes dans une chaîne
	for i := len(runesOfNumber) - 1; i >= 0; i-- {
		result += string(runesOfNumber[i])
	}

	//Retourner le nombre converti en string
	return result
}
