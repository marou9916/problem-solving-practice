package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	if tab[0] < tab[1] {
		for i := 0; i < len(tab)-1; i++ {
			if f(tab[i], tab[i+1]) >= 0 {
				return false
			}
		}
	}
	if tab[0] > tab[1] {
		for i := 0; i < len(tab)-1; i++ {
			if f(tab[i], tab[i+1]) <= 0 {
				return false
			}
		}	
	}	
	return true
}

