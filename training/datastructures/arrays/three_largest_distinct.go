package arrays

func FindThreeLargestDistinctElements(arr []int) []int {
	firstLargest, secondLargest, thirdLargest := -1<<31, -1<<31, -1<<31

	if len(arr) >= 0 && len(arr) < 3 {
		return arr
	}

	for _, element := range arr {
		if element > firstLargest {
			thirdLargest = secondLargest
			secondLargest = firstLargest
			firstLargest = element
		} else if element < firstLargest && element > secondLargest {
			thirdLargest = secondLargest
			secondLargest = element
		} else if element < secondLargest && element > thirdLargest {
			thirdLargest = element
		}
	}

	if thirdLargest == -1<<31 {
		if secondLargest == -1<<31 {
			return []int{firstLargest}
		}
		return []int{firstLargest, secondLargest}
	}

	return []int{firstLargest, secondLargest, thirdLargest}
}
