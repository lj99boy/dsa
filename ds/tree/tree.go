package tree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//func preorderTraversalRecursion(root *TreeNode) []int {
//	holder := make([]int, 0)
//
//	if root.left != nil {
//		root
//	} else {
//		holder = append(holder, root)
//	}
//
//	return holder
//}
