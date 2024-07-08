package training

import "fmt"

func ReverseRange(n1 int, n2 int) {

	if n2-n1 > 0 {
		for i := n2; i >= n1; i-- {
			if i == n1 {
				fmt.Printf("%d", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	}

	if n2-n1 <= 0 {
		for i := n2; i <= n1; i++ {
			if i == n1 {
				fmt.Printf("%d", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Println()
	}
}
