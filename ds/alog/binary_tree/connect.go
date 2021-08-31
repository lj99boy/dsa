package binary_tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	fCombine(root.Left, root.Right)

	return root
}

func fCombine(leftN *Node, rightN *Node) {
	if leftN == nil {
		return
	}

	if leftN.Left == nil {
		leftN.Next = rightN
	} else {
		leftN.Next = rightN
		fCombine(leftN.Left, leftN.Right)
		fCombine(leftN.Right, rightN.Left)
		fCombine(rightN.Left, rightN.Right)
	}
}
