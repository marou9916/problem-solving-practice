package training

func Atoi2(s string) int {
	var x int //Entier à retourner
	decimal := 1
	signe := 1

	//Convertir la string en tableau de runes
	s1 := []rune(s)

	//Parcourir le tableau de runes pour récupérer les digits
	for i := len(s1) - 1; i >= 0; i-- {
		//gérer le signe
		if (s1[i] == '+' || s1[i] == '-') && i == 0 {
			if s1[i] == '-' {
				signe *= -1
				x *= signe
			}
			//stocker les digits
		} else if s1[i] >= '0' && s1[i] <= '9' {
			x += int(s1[i]-'0') * decimal
			decimal *= 10
		} else {
			return 0
		}
	}

	//Retourner le résultat de la conversion
	return x
}
