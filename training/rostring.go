package training

import (
	"fmt"
	"strings"
)

func RoString(input string) {
	var restOftheString []string
	var firstWord = FirstWord(input)

	restOftheString = strings.Split(input, firstWord)

	output := CleanStr(restOftheString[1])

	if len(output) >= 1 {
		fmt.Println(output + " " + firstWord)
	} else {
		fmt.Println(firstWord)
	}
}

// func CleanStr(s1 string) string {
// 	var s2 strings.Builder
// 	temp := ""
// 	addSpace := false

// 	for _, character := range s1 {

// 		if IsSpace(character) {
// 			if temp == "" {
// 				continue
// 			} else {
// 				if addSpace {
// 					s2.WriteString(" ")
// 				}
// 				s2.WriteString(temp)
// 				temp = ""
// 				addSpace = true
// 			}
// 		} else {
// 			temp += string(character)

// 		}
// 	}

// 	if temp != "" {
// 		if addSpace {
// 			s2.WriteString(" ")
// 		}
// 		s2.WriteString(temp)
// 		temp = ""
// 		addSpace = true
// 	}

// 	return s2.String()
// }

// func FirstWord(input string) string {
// 	var output string

// 	for _, character := range input {
// 		if IsSpace(character) {
// 			if output == "" {
// 				continue
// 			} else {
// 				return output
// 			}
// 		} else {
// 			output += string(character)
// 		}
// 	}

// 	return output
// }

// func IsSpace(character rune) bool {
// 	return unicode.IsSpace(character)
// }
