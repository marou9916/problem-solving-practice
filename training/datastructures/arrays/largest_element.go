package arrays

func LargestElement(arr []int) int {
	//Initialiser à zéro la variable de stockage de l'élément le plus grand
	largestElement := arr[0]

	//Rechercher
	//Pour chaque élément du tableau
	for _, element := range arr {
		//Si la valeur d'un élément du tableau est plus grande que la valeur de la variable de stockage
		//ALors affecter cette valeur à la variable
		if largestElement < element {
			largestElement = element
		}
	}

	return largestElement
}
