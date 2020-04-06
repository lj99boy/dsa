package bubblesort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := alog.GenerateSlice(7)
	fmt.Printf("%v", BubbleSort(arr)[:])
}
