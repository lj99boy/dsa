package convertBST

//https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

 type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }

var sum = 0

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sum = 0

	reverse(root)

	return root
}

func reverse(root *TreeNode)  {
	if root.Right != nil {
		reverse(root.Right)
	}

	sum += root.Val
	root.Val = sum

	if root.Left !=nil {
		reverse(root.Left)
	}
}