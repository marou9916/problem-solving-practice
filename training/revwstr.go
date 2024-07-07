package training

import "strings"

func RevWStr(s string) string {
	var builder strings.Builder
	cleanS := CleanStr(s)
	words := strings.Split(cleanS, " ")

	for i := len(words) - 1; i >= 0; i-- {
		if builder.Len() == 0 {
			builder.WriteString(words[i])
		} else {
			builder.WriteString(" ")
			builder.WriteString(words[i])
		}

	}

	return builder.String()
}
