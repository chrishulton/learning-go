package main

import "fmt"

func main() {
	data := []float64{1.0, 1.2, 1.5}

	var sum float64

	for _, float := range data {
		sum += float
	}

	total := float64(len(data))

	fmt.Printf("Average: %f\n", sum/total)
}
