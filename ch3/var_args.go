package main

import "fmt"

func main() {
	varArgs(1, 2)
	varArgs(2, 3, 4, 5, 6)
}

func varArgs(args ...int) {
	for _, num := range args {
		fmt.Printf("%d\n", num)
	}
}
