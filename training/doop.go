package training

import "fmt"

func Doop(x int, operator string, y int) string {
	//Déclaration de la variable de stockage du résultat
	var output int

	//Effectuer l'opération selon l'opérateur
	switch operator {
	case "+":
		output = x + y
	case "-":
		output = x - y
	case "*":
		output = x * y
	case "/":
		if y == 0 {
			fmt.Println("No division by 0")
		} else {
			output = x / y
		}
	case "%":
		if y == 0 {
			fmt.Println("No modulo by 0")
		} else {
			output = x % y
		}
	default:
		break
	}

	// Retourner le résultat de l'opération après conversion en string
	return Itoa(output)
}
