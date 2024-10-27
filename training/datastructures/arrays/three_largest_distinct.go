package arrays

func FindThreeLargestDistinctElements(arr []int) []int {
	if len(arr) >= 0 && len(arr) < 3 { //si l'array contient zéro, un, ou deux éléments alors retourner l'array
		return arr
	}

	//Trouver le premier élément
	firstLargest := LargestElement(arr)

	//Trouver le second élément
	secondLargest := SecondLargestElement(arr)
	if secondLargest == -1 { //s'il n y a pas de second élément distinct plus grand
		return []int{firstLargest} //renvoyer une array contenant uniquement le plus grand élément distinct
	}

	//Trouver le troisième élément
	thirdLargest := ThirdLargestElement(firstLargest, secondLargest, arr)
	if thirdLargest == -1 { //s'il n y a pas de troisième élément distinct plus grand
		return []int{firstLargest, secondLargest} //renvoyer une array contenant les deux plus grands éléments distincts
	}

	return []int{firstLargest, secondLargest, thirdLargest}
}

func ThirdLargestElement(firstLargest, secondLargest int, arr []int) int {
	thirdLargest := -1

	//Parcourir l'array
	for _, element := range arr {
		if element > thirdLargest && element != firstLargest && element != secondLargest {
			thirdLargest = element
		}
	}

	return thirdLargest
}
