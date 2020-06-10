package radixsort

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	//arr := alog.GenerateSlice(7, true)
	arr := []int{71,23,654,467,245,56}

	fmt.Printf("%v", RadixSort(arr)[:])
}
