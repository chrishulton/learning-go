package main

import "fmt"

func main() {
	p := plusTwo()
	fmt.Printf("%v\n", p(2))

	x := plusX(4)
	fmt.Printf("%v\n", x(3))
}

func plusTwo() (f func(int) int) {
	return func(arg int) int {
		return arg + 2
	}
}

func plusX(increment int) (f func(int) int) {
	return func(arg int) int {
		return arg + increment
	}
}
