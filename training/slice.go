package training

func Slice(slice []string, ints ...int) []string {
	start := 0
	end := len(slice)

	if len(ints) == 1 {
		start = ints[0]
		if start < 0 {
			start = len(slice) + start
		}
		if start < 0 || start >= len(slice) {
			return nil // Index hors limites
		}
		return slice[start:end]
	}

	if len(ints) == 2 {
		start, end = ints[0], ints[1]

		if start < 0 {
			start = len(slice) + start
		}
		if end < 0 {
			end = len(slice) + end
		}

		if start < 0 || start >= len(slice) || end < 0 || end > len(slice) || start > end {
			return nil // Index hors limites ou début après la fin
		}
		return slice[start:end]
	}

	return slice[start:end]
}

//Failure

// func Slice(slice []string, ints ...int) []string {
// 	start := 0
// 	end := len(slice)

// 	if len(ints) == 1 {
// 		start = ints[0]

// 		if start > 0 {
// 			return slice[start:end]
// 		} else {
// 			start *= -1
// 			return slice[start-1 : end]
// 		}

// 	} else {
// 		start, end = ints[0], ints[1]

// 		if start > 0 && end > 0 {
// 			if start < end {
// 				return slice[start:end]
// 			} else {
// 				return nil
// 			}
// 		}

// 	}
// 	return slice[start:end]
// }
