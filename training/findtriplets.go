package training

import "sort"

func FindTriplets(arr []int) [][]int {
	var result [][]int

	sort.Ints(arr) //tri du tableau

	for i := 0; i <= len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		sum := arr[i] + arr[left] + arr[right]

		for left < right {
			if sum == 0 {
				result = append(result, []int{i, left, right})

				for left < right && arr[left] == arr[left+1] {
					left++
				}

				for left < right && arr[right] == arr[right-1] {
					right--
				}

				left++
				right--

			} else if sum < 0 {
				left++
				
			} else if sum > 0 {
				right--
			}
		}

	}
	return result
}
