package training

import "math"

func ClosestPair(arr1 []int, arr2 []int, x int) (int, int) {
	var absolutes []int
	positionsInFirstArray := make(map[int]int)
	positionsInSecondArray := make(map[int]int)

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			absolute := int(math.Abs(float64(arr1[i] + arr2[j] - x)))

			absolutes = append(absolutes, absolute)
			positionsInFirstArray[absolute] = i
			positionsInSecondArray[absolute] = j

		}
	}

	minimumAbsolute := getMin(absolutes)

	firstElementPosition := getIndex(positionsInFirstArray, minimumAbsolute)
	secondElementPosition := getIndex(positionsInSecondArray, minimumAbsolute)

	return arr1[firstElementPosition], arr2[secondElementPosition]
}

func getMin(arr []int) int {
	min := arr[0]

	for _, elt := range arr {
		if elt < min {
			min = elt
		}
	}

	return min
}

func getIndex(arr map[int]int, key int) int {
	return arr[key]
}
