package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(4)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
