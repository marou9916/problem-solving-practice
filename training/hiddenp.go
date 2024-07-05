package training

func Hiddenp(s1, s2 string) int {

	if s1 == "" {
		return 1
	}

	index1, index2 := 0, 0

	for index2 <= len(s2)-1 {
		if s2[index2] == s1[index1] {
			index1++
			if index1 == len(s1)-1 {
				return 1
			}
		}
		index2++
	}

	return 0
}
