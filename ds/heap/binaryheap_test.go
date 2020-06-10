package heap

import (
	"fmt"
	"testing"
)

var heap = &BinaryHeap{
	//Holder: []int{1, 3, 5, 8, 9, 20, 31, 44},
	//Holder : []int{2, 3, 7, 9, 18, 25},
	Holder : []int{7, 9, 18, 25},

}

var ruinedHeap = &BinaryHeap{
	Holder: []int{5, 3, 3, 2, 7, 20, 11, 9},
}

//func TestBinaryHeap_AdjustUp(t *testing.T) {
//	heap.AdjustUp(2)
//	fmt.Printf("%v", heap.Holder[:])
//}

func TestBinaryHeap_Pop(t *testing.T) {
	res1 := heap.Pop()
	res2 := heap.Pop()
	fmt.Printf("val:%v,%v,Holder:%v", res1,res2 ,heap.Holder[:])
}

func TestBinaryHeap_Push(t *testing.T) {
	heap.Push(5)
	fmt.Printf("Holder:%v", heap.Holder[:])
}

//func TestBinaryHeap_AdjustDown(t *testing.T) {
//	heap.AdjustDown(0)
//	fmt.Printf("Holder:%v", heap.Holder[:])
//}

func TestBinaryHeap_Build(t *testing.T) {
	ruinedHeap.Build()
	fmt.Printf("Holder:%v", heap.Holder[:])
}
