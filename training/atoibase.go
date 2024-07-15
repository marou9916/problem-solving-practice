package training

import "math"

func AtoiBase(s, b string) int {
	output := 0
	base := 2 //On suppose que la base pour la conversion est binaire
	digits := []int{}

	//Récupérer les digits du nombre
	for i := range s {
		digits = append(digits, int(s[i]-48))
	}

	//Calculer la valeur numérique du nombre dans la base donnée
	for i := len(digits) - 1; i >= 0; i-- {
		exp := len(digits)-1-i
		output += digits[i] * int(math.Pow(float64(base), float64(exp)))
	}

	return output
}
