package stack

import (
	"testing"
)

type chain struct {
	val  int
	next *chain
}

func reverseChain(c *chain) *chain {
	var holder []*chain

	for inChain := c; inChain != nil; inChain = inChain.next {
		holder = append(holder, inChain)
	}

	for i := len(holder) - 1; i > 1; i-- {
		holder[i].next = holder[i-1]
	}

	return holder[len(holder)-1]
}

func Test()  {

}

func TestNew(t *testing.T) {

	//s := New()
	//
	//if !s.isEmpty() ||
	//	s.len != 0 ||
	//	s.Len() != 0 {
	//	t.Error()
	//}
	//
	//s.Push(1)
	//s.Push(2)
	//s.Push(3)
	//
	//if s.stack[0] != 3 ||
	//	s.stack[1] != 2 ||
	//	s.stack[2] != 1 {
	//	fmt.Println(s.stack)
	//	t.Error()
	//}
	//
	//if s.Len() != 3 {
	//	t.Error()
	//}
	//
	//a := s.Pop()
	//
	//if a != 3 || s.Len() != 2 {
	//	t.Error()
	//}
	//
	//b := s.Peek()
	//
	//if b != 2 {
	//	t.Error()
	//}
}
