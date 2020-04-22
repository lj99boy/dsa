package heap

import (
	"fmt"
	"testing"
)

var heap = &BinaryHeap{
	holder: []int{1, 3, 5, 8, 9, 20, 31, 44},
}

var ruinedHeap = &BinaryHeap{
	holder: []int{5, 3, 3, 2, 7, 20, 11, 9},
}

func TestBinaryHeap_AdjustUp(t *testing.T) {
	heap.AdjustUp(2)
	fmt.Printf("%v", heap.holder[:])
}

func TestBinaryHeap_AdjustDown(t *testing.T) {
	heap.AdjustDown(0)
	fmt.Printf("holder:%v", heap.holder[:])
}

func TestBinaryHeap_Build(t *testing.T) {
	ruinedHeap.Build()
	fmt.Printf("holder:%v", heap.holder[:])
}
