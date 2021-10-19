package deleteNode

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

	 //https://leetcode-cn.com/problems/delete-node-in-a-bst/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right,key)
	} else if key < root.Val{
		root.Left = deleteNode(root.Left,key)
	} else {
		if root.Right == nil && root.Left == nil {
			return nil
		}

		if root.Right == nil && root.Left !=nil {
			return root.Left
		}

		if root.Right != nil && root.Left ==nil {
			return root.Right
		}

		// 二叉搜索树 中间某个值要删除的话，要不就是这个值的左边来接替，要不就是右边，左边就表示左子树里最大的，右边就表示右子树里最小的，两种都阔以，提出这个值过后，要在这个子树里把这个值删掉
		if root.Right != nil && root.Left != nil {
			 max := getLeftMax(root.Left)
			 root.Val = max.Val
			 root.Left = deleteNode(root.Left,max.Val)
		}
	}

	return root
}

func getLeftMax(root *TreeNode) *TreeNode  {
	if root.Right == nil {
		return root
	}

	return getLeftMax(root.Right)
}