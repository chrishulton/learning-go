package main

import "fmt"

func main() {
	data := []int{1, 5, 3, 4, 2, 0}
	fmt.Printf("Unsorted: %v\n", data)

	bubbleSort(data)

	fmt.Printf("Sorted: %v\n", data)
}

func bubbleSort(input []int) {
	swapped := true

	for swapped {
		swapped = false

		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				tmp := input[i]
				input[i] = input[i-1]
				input[i-1] = tmp
				swapped = true
			}
		}
	}
}
