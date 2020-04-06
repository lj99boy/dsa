package swap

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	left, right := 5, 3
	left, right = swap(left, right)
	fmt.Printf("%d--%d", left, right)
}
