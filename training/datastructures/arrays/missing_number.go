package arrays

import "errors"

func MissingNumber(arr []int) (int, error) {

	if len(arr) < 1 {
		return -1, errors.New("l'array doit au moins contenir un élément")
	} else if IsThereNegativeNumber(arr) {
		return -1, errors.New("no negative number in the array please")
	}

	numMap := make(map[int]bool) //Créer une map des éléments de l'array

	for _, num := range arr { // Remplir la map(carte)
		numMap[num] = true
	}

	for i := 1; i <= len(arr)+1; i++ { //Vérifier l'élément manquant
		if !numMap[i] {
			return i, nil //Retourner cet élément
		}
	}
	return -1, errors.New("no missing number")
}
func IsThereNegativeNumber(arr []int) bool {
	for _, element := range arr {
		if element < 0 {
			return true
		}
	}
	return false
}
