package mergesort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestMergeSortRecur(t *testing.T) {
	arr := alog.GenerateSlice(7, false)
	xx := MergeSortRecur(arr)[:]
	fmt.Printf("%v", xx)
}
