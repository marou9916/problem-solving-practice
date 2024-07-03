package training

// Split divise la chaîne s en sous-chaînes séparées par le séparateur sep
func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	sLen := len(s)

	if sepLen == 0 {
		for _, char := range s {
			result = append(result, string(char))
		}
		return result
	}

	start := 0
	for i := 0; i <= sLen-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	// Ajouter la dernière sous-chaîne après la dernière occurrence du séparateur
	result = append(result, s[start:])

	return result
}

