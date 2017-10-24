package main

import "fmt"

func main() {
	fmt.Printf("Average: %f\n", floatAvg([]float64{1.0, 1.2, 1.5, 4.6}))
	fmt.Printf("Average: %f\n", floatAvg([]float64{0}))
}

func floatAvg(input []float64) float64 {
	total := float64(len(input))

	if total == 0 {
		return 0
	}

	sum := 0.0

	for _, float := range input {
		sum += float
	}

	return sum / total
}
