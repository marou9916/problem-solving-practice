package training

func ExpandStr(s string) string {
	//Déclarer la string à retourner
	var result string

	//Transformer la string d'entrée en runes
	s1 := []rune(s)

	//Récupérer la chaîne de caractères sans les espaces du début et de la fin
	for i := 0; i <= len(s1)-1; i++ {
		if s1[i] == ' ' && (i == 0 || i == len(s1)-1) { //si un espace à la première et/ou dernière position
			continue
		} else if s1[i] != ' ' { //si pas un espace
			result += string(s1[i])
		} else if s1[i] == ' ' && (i != 0 && i != len(s1)-1) { //si un espace à une position quelconque autre que la première et la dernière
			result += "   " //remplacer par trois espaces
		}
	}

	//Retourner la string
	return result
}
