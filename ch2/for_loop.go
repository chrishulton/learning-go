package main

import "fmt"

func main() {
	_forLoop()
	_goTo()
	_fillArray()
}

func _forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func _goTo() {
	i := 0
Loop:
	if i < 10 {
		fmt.Println(i)
		i++
		goto Loop
	}
}

func _fillArray() {
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}
