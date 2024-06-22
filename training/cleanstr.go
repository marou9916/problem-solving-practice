package training

import (
	"strings"
)

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
	return s2.String()
}
