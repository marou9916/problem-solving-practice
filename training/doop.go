package training

import "fmt"

func Doop(x string, operator string, y string) string {
	//Déclaration de la variable de stockage du résultat
	var output int

	//Convertir en entier les valeurs d'entrées
	value1 := Atoi(x)
	value2 := Atoi(y)

	//Effectuer l'opération selon l'opérateur
	switch operator {
	case "+":
		output = value1 + value2
	case "-":
		output = value1 - value2
	case "*":
		output = value1 * value2
	case "/":
		if value2 == 0 {
			fmt.Println("No division by 0")
		} else {
			output = value1 / value2
		}
	case "%":
		if value2 == 0 {
			fmt.Println("No modulo by 0")
		} else {
			output = value1 % value2
		}
	default:

	}

	//Retourner le résultat de l'opération après conversion en string
	return Itoa(output)
}
