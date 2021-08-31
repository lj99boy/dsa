package find_at_least_max_arr

import (
	"fmt"
	"testing"
)

func TestFindAtLeastMaxArr(t *testing.T) {
	arr := []int{1, 2, 3, 5, 4}
	//arr := []int{1, 2, 4, 3, 5}
	fmt.Printf("%v", findAtLeastMaxArr(arr))
}
