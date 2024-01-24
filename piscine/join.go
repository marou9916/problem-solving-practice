package piscine

func Join(strs []string, sep string) string {
	str := ""
	for index, strTable := range strs {
		str += strTable
		if index != len(strs)-1 {
			str += sep
		}
	}
	return str
}
