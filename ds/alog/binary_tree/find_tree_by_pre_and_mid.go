package binary_tree

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	midIndex := -1
	for k, v := range inorder {
		if v == preorder[0] {
			midIndex = k
		}
	}

	leftS := midIndex - 0

	root.Left = BuildTree(preorder[0+1:0+1+leftS], inorder[0:midIndex])
	if 0+1+leftS+1 > len(preorder) {
		return nil
	}
	root.Right = BuildTree(preorder[0+1+leftS+1:], inorder[midIndex+1:])

	return root
}
