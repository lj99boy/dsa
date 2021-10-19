package isValidBST

//https://leetcode-cn.com/problems/validate-binary-search-tree/

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func isValidBST(root *TreeNode) bool {
return reverse(root,nil,nil)
}

func reverse(root *TreeNode,max *TreeNode,min *TreeNode) bool {
	if root == nil {
		return true
	}

	if max!=nil && root.Val >= max.Val {
		return false
	}

	if min !=nil && root.Val<=min.Val{
		return false
	}

	return reverse(root.Left,root,min) && reverse(root.Right,max,root)
}
