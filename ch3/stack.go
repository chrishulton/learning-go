package main

import "fmt"

var ptr int
var stack []int

func main() {
	stack = make([]int, 10)

	push(3)
	push(6)
	push(8)
	pop()
	push(7)
	push(4)
	pop()

	prettyPrint()
}

func push(i int) {
	stack[ptr] = i
	ptr += 1
}

func pop() (res int) {
	res = stack[ptr]
	ptr -= 1

	return
}

func prettyPrint() {
	for idx, elem := range stack[0:ptr] {
		fmt.Printf("[%d:%d] ", idx, elem)
	}

	fmt.Println()
}
