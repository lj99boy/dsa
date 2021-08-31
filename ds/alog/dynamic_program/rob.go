package dynamic_program

import "math"

func Rob(nums []int) int {

	dTable := make([]int, len(nums)+2)
	var max = 0
	//for i := len(nums) - 1; i >= 0; i-- {
	//	dTable[i] = int(math.Max(float64(dTable[i+2]+nums[i]), float64(dTable[i+1])))
	//	if dTable[i] > max {
	//		max = dTable[i]
	//	}
	//}

	for i := 0; i < len(nums); i++ {
		if i-2
		dTable[i] = int(math.Max(float64(dTable[i+2]+nums[i]), float64(dTable[i+1])))
		if dTable[i] > max {
			max = dTable[i]
		}
	}

	return max
}

func RobCycle(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 0 {
		return 0
	}

	var innerRob = func(nums []int) int {
		dTable := make([]int, len(nums)+2)
		var max = 0
		for i := len(nums) - 1; i >= 0; i-- {
			dTable[i] = int(math.Max(float64(dTable[i+2]+nums[i]), float64(dTable[i+1])))
			if dTable[i] > max {
				max = dTable[i]
			}
		}
		return max
	}

	return int(math.Max(float64(innerRob(nums[0:len(nums)-1])), float64(innerRob(nums[1:]))))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var intMax = 0

func RobTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	dT := make(map[*TreeNode]int)
	var innerMax = 0
	innerRob(root, dT, &innerMax)

	return innerMax
}

func innerRob(root *TreeNode, dT map[*TreeNode]int, innerMax *int) int {
	if v, ok := dT[root]; ok {
		return v
	}

	if root == nil {
		return 0
	}

	var do = root.Val
	if root.Left != nil {
		//do += innerRob(root.Left.Left, dT) + innerRob(root.Left.Right, dT)
		do += innerRob(root.Left.Left, dT, innerMax) + innerRob(root.Left.Right, dT, innerMax)
	}

	if root.Right != nil {
		do += innerRob(root.Right.Left, dT, innerMax) + innerRob(root.Right.Right, dT, innerMax)
	}

	res := int(math.Max(float64(do), float64(innerRob(root.Left, dT, innerMax)+innerRob(root.Right, dT, innerMax))))
	//if res > intMax {
	if res > *innerMax {
		*innerMax = res
		//intMax = res
	}

	dT[root] = res

	return res
}
