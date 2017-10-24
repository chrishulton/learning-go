package main

import "fmt"

func main() {
	data := []int{5, 12, 3, 4, 1, 83, 6}

	fmt.Printf("%d\n", max(data))
}

func max(input []int) int {
	currMax := input[0]

	for _, elem := range input[1:] {
		if elem > currMax {
			currMax = elem
		}
	}

	return currMax
}
