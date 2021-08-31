package tu

import (
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	x:= make([][]int,0)

	x = append(x,[]int{1,2})
	x = append(x,[]int{3})
	x = append(x,[]int{3})
	x = append(x,[]int{})

	allPathsSourceTarget(x)


}
