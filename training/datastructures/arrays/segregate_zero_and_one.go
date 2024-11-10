package arrays

//	func SegregateZeroAndOne(arr []int) {
//		//Trier l'array de 0 et de 1 (bubble sort)
//		for i := 0; i < len(arr); i++ {
//			for j := 0; j < len(arr)-i-1; j++ {
//				if arr[j-1] > arr[j] {
//					arr[j-1], arr[j] = arr[j], arr[j-1]
//				}
//			}
//		}
//	}
func SegregateZeroAndOne(arr []int) {
	type0, type1 := 0, len(arr)-1

	for type0 < type1 {
		if arr[type0] == 0 {
			type0++
		} else if arr[type1] == 1 {
			type1--
		} else {
			arr[type0], arr[type1] = arr[type1], arr[type0]
			type0++
			type1--
		}
	}

}
