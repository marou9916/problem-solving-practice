package training

func Capitalize(s string) string {
	arrays := []rune(s)

	for i := range arrays {
		if arrays[i] == ' ' || arrays[i] == '+' {
			if arrays[i+1] >= 'a' && arrays[i+1] <= 'z' {
				arrays[i+1] = arrays[i+1] - 32
			}
		}
	}

	return string(arrays)
}
