package MinStack

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	xx := Constructor()
	xx.Push(1)
	xx.Push(1)
	xx.Push(2)
	xx.Pop()
	xx.Pop()
	xx.Pop()
	xx.Push(2)
	println(xx.GetMin())
}
