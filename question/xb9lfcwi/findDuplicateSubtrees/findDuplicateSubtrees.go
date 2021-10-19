package findDuplicateSubtrees

import "strconv"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

var holder map[string]int
var res []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	holder = make(map[string]int)
	res = make([]*TreeNode,0)

	reverse(root)

	return res
}

func reverse(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	leftStr := reverse(root.Left)
	rightStr:= reverse(root.Right)

	sRes := leftStr+","+rightStr+","+strconv.Itoa(root.Val)


	if v,ok := holder[sRes];ok && v == 1{
		res = append(res,root)
	}

	holder[sRes]++
	return sRes
}