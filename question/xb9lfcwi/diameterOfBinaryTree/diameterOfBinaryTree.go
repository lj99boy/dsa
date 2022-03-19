package diameterOfBinaryTree

import "math"

//https://leetcode-cn.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	 rec(root)

	return ans
}

var ans int

func rec(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := rec(root.Left)
	rightSum := rec(root.Right)

	tSum := leftSum+rightSum

	if tSum > ans {
		ans = tSum
	}

	return int(math.Max(float64(leftSum), float64(rightSum)) + 1)

}