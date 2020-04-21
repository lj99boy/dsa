package heap

import (
	"fmt"
	"testing"
)

var heap = &BinaryHeap{
	holder: []int{1, 3, 5, 8, 9, 20, 31, 44},
}

func TestBinaryHeap_Push(t *testing.T) {
	heap.Push(2)
	fmt.Printf("%v", heap.holder[:])
}

func TestBinaryHeap_Pop(t *testing.T) {
	res := heap.Pop()
	fmt.Printf("pop:%v -- holder:%v", res, heap.holder[:])
}
