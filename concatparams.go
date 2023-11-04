package piscine

func ConcatParams(args []string) string {
	var result string
	for i := 0; i < len(args); i++ {
		result = result + args[i]
		if i != len(args)-1 {
			result = result + "\n"
		}
	}
	return result
}
