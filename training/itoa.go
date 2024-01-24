package training

func Itoa(the_int int) string {
	//Déclaration des variables
	var runesOfNumber []rune
	number := the_int
	rest := 0
	result := ""

	//Gérer le signe du nombre
	if the_int < 0 {
		result += "-"
		number *= -1
	} else {

	}

	//Récupérer les digits du nombre en runes dans un tableau de runes
	for number > 0 {
		rest, number = number%10, number/10
		runesOfNumber = append(runesOfNumber, rune(rest+'0')) //tableau de runes
	}

	//Concaténer les runes dans une chaîne
	for i := len(runesOfNumber) - 1; i >= 0; i-- {
		result += string(runesOfNumber[i])
	}

	//Retourner le nombre converti en string
	return result
}
