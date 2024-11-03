package arrays

func MissingAndRepeating(arr []int) (int, int) {
	//Initialiser les variables de stockage des résultats
	missing, repeating := -1, -1
	//Controler l'entrée
	if len(arr) < 2 {
		return missing, repeating
	}
	//Trouver le répétitif
	notMissingElements := make(map[int]bool)
	for _, element := range arr {
		if notMissingElements[element] {
			repeating = element
		} else {
			notMissingElements[element] = true
		}
	}

	//Trouver le manquant
	for i := 1; i <= len(arr); i++ {
		if !notMissingElements[i] {
			missing = i
		}
	}
	return missing, repeating
}
