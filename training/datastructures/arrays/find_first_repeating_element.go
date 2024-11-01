package arrays

import "errors"

func FindFirstRepeatingElement(arr []int) (int, error) {

	if len(arr) < 2 {
		return 0, errors.New("pas assez d'éléments dans le tableau")
	}
	elements := make(map[int]bool)

	for _, element := range arr {
		if elements[element] {
			return element, nil
		}
		elements[element] = true
	}

	return 0, errors.New("aucun doublon n'a été trouvé")

}
