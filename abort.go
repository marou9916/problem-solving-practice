package piscine

func Abort(a, b, c, d, e int) int {
	nums := [5]int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	middle := nums[2]
	return middle
}
