package shellsort

import (
	"dsa/ds/alog"
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	arr := alog.GenerateSlice(8,false)

	fmt.Printf("%v", ShellSort(arr)[:])
}
