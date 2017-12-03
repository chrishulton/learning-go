package main

import "fmt"

type Maxer interface {
	Get(i int) interface{}
	Length() int
	Less(i, j interface{}) bool
}

func Max(x Maxer) (max interface{}) {
	max = x.Get(0)

	for i := 0; i < x.Length(); i++ {
		elem := x.Get(i)

		if x.Less(max, elem) {
			max = elem
		}
	}

	return
}

type Xi []int

func (p Xi) Get(i int) interface{} { return p[i] }
func (p Xi) Length() int           { return len(p) }
func (p Xi) Less(i, j interface{}) bool {
	// lousy, but this doesn't seem well suited to generics anyway
	switch i.(type) {
	case int:
		if _, ok := j.(int); ok {
			return i.(int) < j.(int)
		} else {
			panic("Expected second arg to be int")
		}
	default:
		panic("Expected first arg to be int")
	}
}

type Xf []float32

func (p Xf) Get(i int) interface{} { return p[i] }
func (p Xf) Length() int           { return len(p) }
func (p Xf) Less(i, j interface{}) bool {
	switch i.(type) {
	case float32:
		if _, ok := j.(float32); ok {
			return i.(float32) < j.(float32)
		} else {
			panic("Expected second arg to be float32")
		}
	default:
		panic("Expected first arg to be float32")
	}
}

func main() {
	ints := Xi{5, 12, 3, 4, 1, 83, 6}
	fmt.Printf("%v\n", Max(ints))

	floats := Xf{5.01, 12.83, 12.56, 4.21, 1.8}
	fmt.Printf("%v\n", Max(floats))
}
