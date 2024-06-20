package training

import (
	"strconv"
	"sort"
	"fmt"
)

func Rn(input string) string {
	var romanNumber string

	romanNumerals := map[int]string {
		1 : "I",
		4 : "IV",
		5 : "V",
		9 : "IX",
		10: "X",
		40: "XL",
		50: "L",
		90: "XC",
		100:"C",
		400:"CD",
		500:"D",
		900:"CM",
		1000:"M",
	}

	//Convertir l'input en nombre entier
	number, err := strconv.Atoi(input)
	if err != nil || number < 1 || number >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return ""
	}

	//Si le nombre est présent dans la map, renvoyer directement la valeur romaine associée
	if IsInTheMap(romanNumerals, number) {
		return romanNumerals[number]
	} 

	//Sinon décomposer le nombre
	//Récupérer les clés de la map
	keys := (make([]int,0, len(romanNumerals)))
	for key := range romanNumerals {
		keys = append(keys, key)
	}

	//Trier dans l'ordre décroissant
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	//Parcourir les clés triées
	for _, key := range keys {
		for number >= key {
			romanNumber += romanNumerals[key]
			number -= key
		}
	}

	return romanNumber
}

func IsInTheMap(romanNumerals map[int]string, number int) bool {
	for key := range romanNumerals {
		if key == number {
			return true
		}
	}
	return false
}