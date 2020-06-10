package binarysearch

import (
	"dsa/ds/alog"
	"dsa/ds/alog/sort/quicksort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := quicksort.QuickSort(alog.GenerateSlice(9, false))
	needle := data[0]

	if BinarySearch(needle, data) != 0 {
		t.Error("res wrong")
	}

}
