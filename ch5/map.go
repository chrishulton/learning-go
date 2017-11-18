package main

import (
	"fmt"
	"strings"
)

type T interface{}

func myMap(input []T, fn func(T) T) []T {
	result := make([]T, len(input))

	for idx, elem := range input {
		result[idx] = fn(elem)
	}

	return result
}

func mapAndPrint(input []T, f func(i T) T) {
	for _, mapped := range myMap(input, f) {
		fmt.Printf("%v\n", mapped)
	}
}

func main() {
	intData := []T{1, 2, 3, 4}
	// chapter should have explained type assertions!
	intFunc := func(i T) T { return i.(int) * 3 }

	mapAndPrint(intData, intFunc)

	strData := []T{"these", "are", "some", "strings"}
	strFunc := func(i T) T { return strings.ToUpper(i.(string)) }

	mapAndPrint(strData, strFunc)
}
