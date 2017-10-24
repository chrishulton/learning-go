package main

import "fmt"

func main() {
	for _, val := range fibonacci(12) {
		fmt.Printf("%d\n", val)
	}
}

func fibonacci(terms int) []int {
	results := make([]int, terms)

	for i := 0; i < terms; i++ {
		if i < 2 {
			results[i] = 1
		} else {
			results[i] = results[i-1] + results[i-2]
		}
	}

	return results
}
