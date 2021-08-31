package dynamic_program

import "testing"

func TestMaxProfitOneTrans(t *testing.T) {
	//tArr := []int{7, 1, 5, 3, 6, 4}
	tArr := []int{7, 6, 4, 3, 1}
	println(MaxProfitOneTrans(tArr))
}

func TestMaxProfitDesiredTimes(t *testing.T) {
	tArr := []int{3, 3, 5, 0, 0, 3, 1, 4}
	//tArr := []int{1, 2, 3, 4, 5}
	println(MaxProfitDesiredTimes(tArr))
}

func TestMaxProfitAnyTimesWithCooldown(t *testing.T) {
	tArr := []int{1, 2, 3, 0, 2}
	println(MaxProfitAnyTimesWithCooldown(tArr))
}

func TestRob(t *testing.T) {
	//tArr := []int{1, 2, 3, 1}
	tArr := []int{1, 2, 3, 1}
	println(Rob(tArr))
}

func TestRobTree(t *testing.T) {
	tTree := &TreeNode{Left:&TreeNode{Val:2},Right:&TreeNode{Val:3},Val:1}

	println(RobTree(tTree))
}