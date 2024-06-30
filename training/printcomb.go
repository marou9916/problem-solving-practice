package training

import (
	"fmt"
	"strings"
)

// 2nde version
func PrintComb() {
	var combinations []string

	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				combinations = append(combinations, fmt.Sprintf("%d%d%d", i, j, k))
			}
		}
	}

	fmt.Println(strings.Join(combinations, ", "))
}

//1ere version

// func PrintComb() {
// for i := 0; i <= 7; i++ {
// 	for j := 1; j <= 8; j++ {
// 		for k := 2; k <= 9; k++ {
// 				if i == 7 && j == 8 && k == 9 {
// 					os.Stdout.WriteString(Itoa(i) + Itoa(j) + Itoa(k))
// 					break
// 				}
//  				os.Stdout.WriteString(Itoa(i) + Itoa(j) + Itoa(k) + ", ")
// 			}
// 		}
// 	}

// 	os.Stdout.WriteString("\n")
// }
