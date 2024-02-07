package training

func Capitalize(s string) string {
	var result string //la variable à retourner après remplacement des lettres minuscules, en début de mot, en lettres majuscules
	s2 := []rune(s)   //conversion de la string en runes

	//Transformer en majuscules les minuscules de chaque début de mot
	for i := 0; i <= len(s2)-1; i++ {
		if s2[i] >= 'a' && s2[i] <= 'z' {
			if i == 0 || s2[i-1] == ' ' {
				s2[i] = s2[i] - 32
			}
		}
	}

	//Revenir en string
	for j := 0; j <= len(s2)-1; j++ {
		result += string(s2[j])
	}

	//Retourner la nouvelle string
	return result
}
