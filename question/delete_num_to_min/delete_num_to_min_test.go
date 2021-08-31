package delete_num_to_min

import (
	"fmt"
	"testing"
)

func TestDeleteNumToMin(t *testing.T) {
	//arr := []int{1, 5, 9, 3, 2, 1, 2}
	arr := []int{3,0,2,0,0}
	//fmt.Printf("%v",deleteNumToMin(arr,3))
	fmt.Printf("%v",deleteNumToMin(arr,1))
}
