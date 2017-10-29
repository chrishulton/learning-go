package main

import (
	"bufio"
	"fmt"
	"github.com/chrishulton/learning-go/ch4/stack"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter balanced operators and operands (ex `3 14 + 5 *`): ")

	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", 1)

	tokens := strings.Split(input, " ")
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
		} else if token == "+" {
			stack.Push(stack.Pop() + stack.Pop())
		} else if token == "-" {
			tail := stack.Pop()
			head := stack.Pop()
			stack.Push(head - tail)
		} else if token == "*" {
			stack.Push(stack.Pop() * stack.Pop())
		} else if token == "/" {
			tail := stack.Pop()
			head := stack.Pop()
			stack.Push(head / tail)
		} else {
			panic("Got unexpected input")
		}
	}

	fmt.Printf("Result is: %d\n", stack.Pop())
}
