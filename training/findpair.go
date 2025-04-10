package training

import (
	"math"
	"sort"
)

func FindPair(arr []int, x uint) string {
	//Sort array
	sort.Ints(arr)

	left, right := 0, len(arr)-1

	for left < right {
		if math.Abs(float64(arr[left])-float64(arr[right])) == float64(x) {
			return "Yes"
		} else if math.Abs(float64(arr[left])-float64(arr[right])) > float64(x) {
			right--
		} else if math.Abs(float64(arr[left])-float64(arr[right])) < float64(x) {
			left++
		}
	}
	return "No"
}
