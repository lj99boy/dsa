package binarysearch

import (
	"dsa/ds/alog"
	"dsa/ds/alog/quicksort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := quicksort.QuickSort(alog.GenerateSlice(9))
	needle := data[3]

	if BinarySearch(needle, data) != 3 {
		t.Error("res wrong")
	}

}
