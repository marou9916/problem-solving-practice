package arrays

func SelectionSort() {
	arr := []int{48, 38, 34, 19, 49, 9, 14, 16, 26}

	for i := 0; i < len(arr); i++ {
		minimumUnSorted := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < minimumUnSorted {
				minimumUnSorted = j
			}
		}
		arr[i], arr[minimumUnSorted] = arr[minimumUnSorted], arr[i]
	}

}
