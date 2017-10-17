package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fizz := (i % 3) == 0
		buzz := (i % 5) == 0

		if fizz {
			fmt.Print("Fizz")
		}

		if buzz {
			fmt.Print("Buzz")
		}

		if !fizz && !buzz {
			fmt.Print(i)
		}

		fmt.Print("\n")
	}
}
