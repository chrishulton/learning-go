package main

import "fmt"

type T interface{}

type Node struct {
	value T
	next  *Node
}

type List struct {
	head *Node
}

func (list *List) addToFront(data T) {
	head := list.head

	list.head = &Node{value: data, next: head}
}

func main() {
	var list List
	list.addToFront(1)
	list.addToFront(2)
	list.addToFront(4)

	for n := list.head; n != nil; n = n.next {
		fmt.Printf("%v\n", n.value)
	}
}
