package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	result := [][]int{}
	current := []int{}
	result = GenerateSubSet(arr, 0, current, result)

	for _, subset := range result {
		fmt.Println(subset)
	}
}

func GenerateSubSet(arr []int, index int, current []int, result [][]int) [][]int {
	temp := make([]int, len(current))
	copy(temp, current)
	result = append(result, temp)

	for i := index; i < len(arr); i++ {
		current = append(current, arr[i])

		result = GenerateSubSet(arr, i+1, current, result)

		current = current[:len(current)-1]
	}
	return result
}
