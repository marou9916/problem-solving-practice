package training

func LinearSearch(input []int, target int) int {
	for index, value := range input {
		if value == target {
			return index
		}
	}

	return -1
}



