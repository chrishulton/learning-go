package stack

import "testing"

func TestPushAndPop(t *testing.T) {
	Reset()
	Push(7)

	if Pop() != 7 {
		t.Log("Popped value should equal pushed value")
		t.Fail()
	}
}

func TestMultiPushAndPop(t *testing.T) {
	Reset()
	Push(1)
	Push(2)
	Push(3)
	Pop()
	Pop()
	Push(2)
	Push(3)
	Push(4)
	Push(5)

	if Pop() != 5 {
		t.Log("Popped value should equal pushed value")
		t.Fail()
	}
}

func TestPopWithoutElements(t *testing.T) {
	defer func() {
		if x := recover(); x == nil {
			t.Log("Should error when popping with no elements")
			t.Fail()
		}
	}()

	Reset()
	Pop()
}
