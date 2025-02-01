package training

import "fmt"

func FindTriplets(arr []int) [][]int {
	var result [][]int

	for i := 0; i <= len(arr)-1; i++ { //prendre un 1er élément à une position
		for j := i + 1; j <= len(arr)-1; j++ { //prendre un 2nd élément à une autre position
			if j == i {
				continue
			}
			for k := j + 1; k <= len(arr)-1; k++ { //prendre un 3eme élément à une autre position
				if k == j {
					continue //sauter les cas de doublons
				}
				fmt.Printf("Triplet{%d, %d, %d},\n", i, j, k)
				if arr[i] + arr[j] + arr[k] == 0 { //si la somme est zéro
					fmt.Printf("Solution ajoutée{%d, %d, %d},\n", i, j, k) //enregistrer le triplet 
					result = append(result, []int{i, j, k})
				}
			}
		}
	}
	return result
}
