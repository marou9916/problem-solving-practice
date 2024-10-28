package arrays

import "errors"

// First version
func MissingNumber(arr []int) (int, error) {
	missingOne := 0

	if len(arr) < 1 && NoNegativeNumber(arr) {
		for i := 1; i <= len(arr)+1; i++ {
			if !ElementIsNotInArray(arr, i) {
				continue
			} else {
				return i, nil
			}
		}
	} else {
		return -1, errors.New("no negative number in array")
	}
	return missingOne, nil
}

func ElementIsNotInArray(arr []int, element int) bool {
	for index, x := range arr {
		if element == x {
			return false
		} else if index == len(arr)-1 {
			return true
		}
	}
	return false
}

func NoNegativeNumber(arr []int) bool {
	for _, element := range arr {
		if element > 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

//Second one
