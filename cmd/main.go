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
	x := args[0]
	operator := args[1]
	y := args[2]

	//Afficher le résultat de l'opération
	fmt.Println(training.Doop(x, operator, y))
}
