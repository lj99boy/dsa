package minPathSum

import "testing"

func Test_minPathSum(t *testing.T) {
	tG := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	tG2 := [][]int{
		{1,2,3},
		{4,5,6},
	}

	println(minPathSum(tG))
	println(minPathSum(tG2))
}
