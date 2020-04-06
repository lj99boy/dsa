package binarysearch

import (
	"dsa/ds/alog"
	"dsa/ds/alog/sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := sort.QuickSort(alog.GenerateSlice(9))
	needle := data[0]

	if BinarySearch(needle, data) != 0 {
		t.Error("res wrong")
	}

}
