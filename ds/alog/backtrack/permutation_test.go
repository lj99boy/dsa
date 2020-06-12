package backtrack

import (
	"fmt"
	"testing"
)

func TestMainGetPermutation(t *testing.T) {
	in := []int{1, 2, 3}

	res := MainGetPermutation(in)

	for _, val := range res {
		fmt.Printf("%v", val[:])
	}
}
