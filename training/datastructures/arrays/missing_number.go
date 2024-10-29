package arrays

import "errors"

func MissingNumber(arr []int) (int, error) {

	if len(arr) < 1 {
		return -1, errors.New("l'array doit au moins contenir un élément")
	} else if IsThereNegativeNumber(arr) {
		return -1, errors.New("no negative number in the array please")
	}
	n := len(arr) + 1

	//CAlculer la somme des entiers de l'array
	sum1 := 0
	for _, element := range arr {
		sum1 += element
	}

	//CAlculer la somme des n premiers entiers
	sum2 := n * (n + 1) / 2

	//Faire la différence des sommes pour obtenir l'élément manquant
	return (sum2 - sum1), nil
}
func IsThereNegativeNumber(arr []int) bool {
	for _, element := range arr {
		if element < 0 {
			return true
		}
	}
	return false
}
