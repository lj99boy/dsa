package countNodes

import "math"
//https://leetcode-cn.com/problems/count-complete-tree-nodes/submissions/
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l:=root
	r:=root

	lH:=0
	for ;l.Left!=nil;{
		lH+=1
		l= l.Left
	}
	rH:=0
	for ;r.Right!=nil;{
		rH+=1
		r = r.Right
	}

	if lH == rH {
		return int(math.Pow(2, float64(lH+1)))-1
	}

	return 1+countNodes(root.Left)+countNodes(root.Right)
}
