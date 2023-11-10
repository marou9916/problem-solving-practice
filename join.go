package piscine

func Join(strs []string, sep string) string {
	concatenated := ""
	for index, strTable := range strs {
		concatenated += strTable
		if index != len(concatenated)-1 {
			concatenated += sep
		}
	}
	return concatenated
}
