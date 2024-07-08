package training

// Swap échange les éléments pointés par x1 et x2
func Swap(x1, x2 *string) {
	*x1, *x2 = *x2, *x1
}

// SortWordArr trie le tableau elements en utilisant un tri à bulles complet
func SortWordArr(elements []string) {
	n := len(elements)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if elements[j] > elements[j+1] {
				Swap(&elements[j], &elements[j+1])
			}
		}
	}
}
