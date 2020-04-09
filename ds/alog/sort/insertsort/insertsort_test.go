package insertsort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := alog.GenerateSlice(7, false)
	fmt.Printf("%v", InsertSort(arr[:]))
}
