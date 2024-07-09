package training

import "fmt"

func Range(n1 int, n2 int) {

	if n1-n2 > 0 {
		for i := n1; i >= n2; i-- {
			if i == n2 {
				fmt.Printf("%d", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	}

	if n1-n2 <= 0 {
		for i := n1; i <= n2; i++ {
			if i == n2 {
				fmt.Printf("%d", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	}
}
