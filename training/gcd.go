package training

func Gcd(a, b uint) uint {
	// Si l'un des nombres est 0, retourner 0
	if a == 0 || b == 0 {
		return 0
	}

	// Utilisation de l'algorithme d'Euclide pour trouver le PGCD
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
