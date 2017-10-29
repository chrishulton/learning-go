package stack

var ptr int
var slice []int

// waiting for the 'Objects' chapter...
func Reset() {
	ptr = 0
	slice = []int{}
}

func Push(i int) {
	newLength := ptr + 1

	if len(slice) < newLength {
		bigger := make([]int, (len(slice)*2)+1)
		copy(bigger, slice[0:])
		slice = bigger
	}

	slice[ptr] = i
	ptr = newLength
}

func Pop() (res int) {
	if ptr > 0 {
		ptr -= 1
		res = slice[ptr]
	} else {
		panic("No elements left to pop!")
	}

	return
}
