package arrays

func InsertionSort() {
	arr := []int{48, 38, 34, 19, 49, 9, 14, 16, 26}

	for i := 1; i < len(arr); i++ {
		extractedElement := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > extractedElement {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = extractedElement
	}
}
