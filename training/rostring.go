package training

import (
	"fmt"
	"strings"
)

func RoString(input string) {
	var restOftheString []string
	var firstWord = FirstWord(input)

	fmt.Println(firstWord)
	

	restOftheString = strings.Split(input, firstWord)
	fmt.Println(restOftheString[1])
	fmt.Println(CleanStr(restOftheString[1]) + " " + firstWord)




}
