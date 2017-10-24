package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		printNum(i)
	}
}

func printNum(num int) {
	fmt.Printf("%d\n", num)
}
