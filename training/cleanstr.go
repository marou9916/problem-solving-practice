package training

import "strings"

func CleanStr(s1 string) string {
	var s2 strings.Builder
	temp := ""
	addSpace := false

	for _, character := range s1 {

		if IsSpace(character) {
			if temp == "" {
				continue
			} else {
				if addSpace {
					s2.WriteString(" ")
				}
				s2.WriteString(temp)
				temp = ""
				addSpace = true
			}
		} else {
			temp += string(character)

		}
	}

	if temp != "" {
		if addSpace {
			s2.WriteString(" ")
		}
		s2.WriteString(temp)
		temp = ""
		addSpace = true
	}

	return s2.String()
}

// func CleanStr(s1 *string) {
// 	var builder strings.Builder
// 	inWord := false

// 	for _, char := range *s1 {
// 		if IsSpace(char) {
// 			inWord = false
// 		} else {
// 			if builder.Len() > 0 && !inWord {
// 				builder.WriteRune(' ')
// 			}
// 			builder.WriteRune(char)
// 			inWord = true
// 		}
// 	}

// 	*s1 = builder.String()
// }
