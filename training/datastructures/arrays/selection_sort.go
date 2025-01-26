package arrays

func SelectionSort() {
	arr := []int{48, 38, 34, 19, 49, 9, 14, 16, 26}

	for i := 0; i < len(arr); i++ {
		minIndx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndx] {
				minIndx = j
			}
		}
		arr[minIndx], arr[i] = arr[i], arr[minIndx]
	}
}
