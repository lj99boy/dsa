package binary_tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	if root.Right == nil {
		root.Right = tmp
		return
	}
	var i *TreeNode
	for i = root.Right; i.Right != nil; i = i.Right {
	}

	i.Right = tmp
}
