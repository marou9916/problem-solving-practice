package training

import "fmt"

func ExpandStr(s string) {
	var temp string
	var words []string

	//Stocker les mots
	for _, char := range s {
		if char != ' ' {
			temp += string(char)
		} else if temp != "" {
			words = append(words, temp)
			temp = ""
		}
	}

	words = append(words, temp)

	//Afficher les mots avec trois esapces entre chaque
	for i, word := range words {
		fmt.Print(word)
		if i != len(words)-1 {
			fmt.Print("   ")
		}
	}
	fmt.Println()

}
