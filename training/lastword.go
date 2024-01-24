package training

import (
	"fmt"
	"strings"
)

func LastWord(the_string string) {
	words_of_the_string := strings.Split(the_string, " ")
	if len(words_of_the_string) == 1 {
		fmt.Println(words_of_the_string[0])
	} else if len(words_of_the_string) > 1 {
		fmt.Println(words_of_the_string[len(words_of_the_string)-1])
	}
}
