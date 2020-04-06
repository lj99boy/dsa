package combsort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestCombSort(t *testing.T) {
	arr := alog.GenerateSlice(7)
	fmt.Printf("%v", CombSort(arr)[:])
}
