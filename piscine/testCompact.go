package piscine

func TestCompact(ptr *[]string) int {
	var compactedList []string
	for _, x := range *ptr {
		if x != "" {
			compactedList = append(compactedList, x)
		}
	}
	*ptr = compactedList
	return len(compactedList)
}
