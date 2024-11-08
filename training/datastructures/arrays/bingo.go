package arrays

import (
	"time"

	"golang.org/x/exp/rand"
)

func Bingo(arr [5][5]int) bool {
	//Tirer un numéro aléatoire
	rand.Seed(uint64(time.Now().UnixNano())) //Nécessaire pour générer un numéro différent à chaque exécution du programme
	randomNumber := rand.Intn(100)

	//Parcourir la grille par colonnes
	counter := 0 //compteur de jetons

	for column := 0; column < 5; column++ {
		for line := 0; line < 5; line++ {
			//Vérifier  si la valeur de la cellule est égale à celle du nombre aléatoire
			if arr[line][column] == randomNumber {
				arr[line][column] = -1
				counter++
				if counter == 5 { //Si la colonne entière est marquée
					return true
				}
			} else {
				if column == 4 { //Si la dernière colonne est atteinte
					return false
				}
				continue //Si ce n'est pas la dernière colonne, on  passe à la suivante
			}
		}
	}
	return false
}
