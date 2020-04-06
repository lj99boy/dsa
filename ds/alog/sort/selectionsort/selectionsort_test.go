package selectionsort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := alog.GenerateSlice(7)
	fmt.Printf("%v\n", arr[:])
	fmt.Printf("%v\n", SelectionSort(arr)[:])
}
