package arrays

func SecondLargestElement(arr []int) int {
	largest, secondLargest := -1, -1

	for _, element := range arr {
		if element > largest {
			secondLargest = largest
			largest = element
 		} else if element < largest && element > secondLargest {
			secondLargest = element
		}
	}

	return secondLargest
}