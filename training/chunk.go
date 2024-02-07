package training

import "fmt"

func Chunk(slice []int, size int) {
	var j int
	if size != 0 {
		for i := 0; i <= len(slice)-1; i += size {
			j += size
			if j > len(slice) {
				j = len(slice)
			}
			fmt.Print(slice[i:j])
		}
		fmt.Println()
	} else {
		fmt.Println()
	}

}
