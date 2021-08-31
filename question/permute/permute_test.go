package permute

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	//tChoice := []int{1,2,3}
	tChoice := []int{5,4,6,2}
	//tChoice := []int{0,1}
	fmt.Printf("%v",Permute(tChoice))
}
