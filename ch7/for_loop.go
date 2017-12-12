package main

import "fmt"

var ci chan int
var cb chan bool

func main() {
	ci = make(chan int)
	cb = make(chan bool)

	go printNums()

	for i := 0; i < 10; i++ {
		ci <- i
	}

	cb <- true
}

func printNums() {
	for {
		select {
		case num := <-ci:
			fmt.Printf("%d\n", num)
		case <-cb:
			break
		}
	}
}
