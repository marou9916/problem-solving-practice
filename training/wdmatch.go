package training

func WdMatch(string1, string2 string) bool {
	index1, index2 := 0, 0

	for index2 < len(string2) {
		if string1[index1] == string2[index2] {
			index1++
			if index1 == len(string1) {
				return true
			}
		}
		index2++
	}

	return false
}
