package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4}
	f := func(i int) int { return i * 3 }

	for _, mapped := range myMap(data, f) {
		fmt.Printf("%d\n", mapped)
	}
}

func myMap(input []int, fn func(int) int) []int {
	result := make([]int, len(input))

	for idx, elem := range input {
		result[idx] = fn(elem)
	}

	return result
}
