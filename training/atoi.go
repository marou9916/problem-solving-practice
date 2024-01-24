package training

func Atoi(a_string string) int {
	//Convertir la chaîne en runes
	the_string_in_runes := []rune(a_string)

	//Initialisation des variables
	the_string_in_int := 0
	signe := 1
	decimal := 1

	//Parcours de la chaîne en sens inverse
	for i := len(the_string_in_runes) - 1; i >= 0; i-- {
		//Vérifier les caractères invalides (espaces ou lettres)
		if the_string_in_runes[i] == ' ' || (the_string_in_runes[i] >= 'A' && the_string_in_runes[i] <= 'Z') || (the_string_in_runes[i] >= 'a' && the_string_in_runes[i] <= 'z') {
			return 0
		} else if the_string_in_runes[i] == '-' || the_string_in_runes[i] == '+' {
			//Gestion des signes
			if the_string_in_runes[i+1] == '-' || the_string_in_runes[i+1] == '+' {
				return 0
			} else if the_string_in_runes[i] == '-' {
				signe *= -1
				the_string_in_int *= signe
			} else if the_string_in_runes[i] == '+' {
				the_string_in_int *= signe
			}

		} else {
			//Conversion du chiffre en entier
			the_string_in_int += int(the_string_in_runes[i]-'0') * decimal
			decimal *= 10
		}
	}

	//Retourner la valeur convertie
	return the_string_in_int
}
