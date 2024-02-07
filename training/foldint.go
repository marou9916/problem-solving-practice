package training

import "fmt"

func FoldInt(f func(int, int) int, table []int, n int) {
	result := f(n, table[0])
	for i := 1; i <= len(table)-1; i++ {
		result = f(result, table[i])
	}

	fmt.Println(result)
}
