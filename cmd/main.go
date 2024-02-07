package main

import (
	"fmt"
	"os"
	"piscine-go/training"
)

func main() {
	//Récupérer les arguments rentrés en ligne de commande
	args := os.Args[1:]

	//Contrôler la saisie
	if len(args) != 3 || args[1] != "+" && args[1] != "-" && args[1] != "*" && args[1] != "/" && args[1] != "%" {
		return
	}

	//Récupérer la valeur1, l'opérateur et la valeur2
	x := training.Atoi2(args[0]) //conversion de la valeur1
	operator := args[1]
	y := training.Atoi2(args[2]) //conversion de la valeur2
	result := training.Doop(x, operator, y)

	//Si les valeurs d'entrées sont comprises entre les valeurs seuils on affiche le résultat de l'opération
	if (x >= -2147483647 && x <= 2147483646) && (y >= -2147483647 && y <= 2147483646) {
		if x == 0 && args[0] != "0" || y == 0 && args[2] != "0" {
			return
		}
		fmt.Println(result)
	}

}
