package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{4, 6, 2, 7, 1, 9, 0, 4}
	data = QuickSort(data)
	fmt.Printf("%#v", data[:])
}

func TestQuickSortSingleMap(t *testing.T) {
	data := []int{4, 6, 2, 7, 1, 9, 0, 4}
	data = QuickSortSingleMap(data)
	fmt.Printf("%#v", data[:])
}
