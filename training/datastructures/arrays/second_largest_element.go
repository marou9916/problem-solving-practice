package arrays

func SecondLargestElement(arr []int) int {

	n := len(arr)

	if n < 2 {
		return -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; n++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}   
	}
	
	largest := arr[0]

	for _, element := range arr {
		if element < largest {
			return element
		}
	}


	return -1
}
