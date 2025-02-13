package training

// func PartitionPoint(arr []int) int {
// 	var result int = -1

// 	for i := 1; i <= len(arr)-2; i++ {
// 		currentElement := arr[i]
// 		left, right := i-1, i+1

// 		for left >= 0 && right <= len(arr)-1 {
// 			if arr[left] < currentElement {
// 				left--
// 			} else {
// 				return -1
// 			}

// 			if arr[right] > currentElement {
// 				right++
// 			} else {
// 				return -1
// 			}

// 			if left == 0 && right == len(arr)-1 {
// 				return i
// 			}
// 		}
// 	}
// 	return result
// }

func PartitionPoint(arr []int) int {
	var result int = -1

	for i := 1; i <= len(arr)-2; i++ {
		currentElement := arr[i]
		maxLeft, minRight := arr[i-1], arr[i+1]

		//Trouver le max à gauche
		for j := i - 1; j >= 0; j-- {
			if arr[j] > maxLeft {
				maxLeft = arr[j]
			}
		}

		//Trouver le min à droite
		for k := i + 1; k <= len(arr)-1; k++ {
			if arr[k] < minRight {
				minRight = arr[k]
			}
		}

		for  maxLeft < minRight {
			if currentElement > maxLeft && currentElement < minRight {
				return i
			}
		}

	}
	return result
}
