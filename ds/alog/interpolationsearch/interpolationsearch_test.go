package interpolationsearch

import (
	"dsa/ds/alog"
	"dsa/ds/alog/sort"
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	data := sort.QuickSort(alog.GenerateSlice(9))

	if InterpolationSearch(data, data[0], 0, len(data)-1) != 0 {
		t.Error("wrong error")
	}
}
