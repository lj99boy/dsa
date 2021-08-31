package searchRange

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tInt := []int{5, 7, 7, 8, 8, 10}
	//tInt := []int{2,2}
	fmt.Printf("%v",SearchRange(tInt,8))
	//fmt.Printf("%v",SearchRange(tInt,6))
}
