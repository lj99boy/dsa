package other

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{0, 0, 0}

	fmt.Printf("%v", ThreeSum(nums))
}

func TestFourSum(t *testing.T) {
	//nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{2,2,2,2,2}


	fmt.Printf("%v", FourSum(nums,8))
}
