package binary_tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	root := &TreeNode{}
	tmpB := -1
	tmpK := -1
	for k, v := range nums {
		if v > tmpB {
			tmpB = v
			tmpK = k
		}
	}
	root.Val = tmpB
	root.Left = constructMaximumBinaryTree(nums[:tmpK])
	root.Right = constructMaximumBinaryTree(nums[tmpK+1:])

	return root
}
