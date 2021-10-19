package searchBST

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

//https://leetcode-cn.com/problems/search-in-a-binary-search-tree/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}


	if val == root.Val {
		return root
	} else if val>root.Val {
		return searchBST(root.Right,val)
	} else {
		return searchBST(root.Left,val)
	}
}
