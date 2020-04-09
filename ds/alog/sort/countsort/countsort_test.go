package countsort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	arr := alog.GenerateSlice(7, false)

	fmt.Printf("%v", CountSort(arr)[:])
}
