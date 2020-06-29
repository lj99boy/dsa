package pancake

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	arr := []int{3, 2, 5, 7, 2, 1}
	Basic(arr)
	fmt.Printf("%v", arr)
}
