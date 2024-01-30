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
	if len(args) != 3 {
		return
	}

	//Récupérer la valeur1, l'opérateur et la valeur2
	x := training.Atoi2(args[0]) //conversion de la valeur1
	operator := args[1]
	y := training.Atoi2(args[2]) //conversion de la valeur2

	//Si les valeurs d'entrées dépassent les valeurs seuils on ne renvoie rien
	if (x <= -2147483648 || x >= 2147483647) || (y <= -2147483648 || y >= 2147483647) {
		return
	}

	//Afficher le résultat de l'opération
	fmt.Print(training.Doop(x, operator, y))
}
