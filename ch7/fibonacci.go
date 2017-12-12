package main

import "fmt"

var ci chan int
var cb chan bool

func main() {
	ci = make(chan int)
	cb = make(chan bool)

	go func() {
		for i := 1; i < 10; i++ {
			// block until term is available
			fmt.Println(<-ci)
		}

		cb <- true
	}()

	fibonacci()
}

func fibonacci() {
	a, b := 1, 1

	for {
		select {
		case ci <- a:
			a, b = b, a+b
		case <-cb:
			return
		}
	}
}
