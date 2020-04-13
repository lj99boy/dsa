package tree

import "testing"

var tree = &ChainBinaryTree{
	Val: 1,
	Left: &ChainBinaryTree{
		Val: 2,
		Left: &ChainBinaryTree{
			Val: 4,
		},
		Right: &ChainBinaryTree{
			Val: 5,
		},
	},
	Right: &ChainBinaryTree{
		Val: 3,
		Left: &ChainBinaryTree{
			Val: 6,
		},
		Right: &ChainBinaryTree{
			Val: 7,
		},
	},
}

func TestPreOrderTraversal(t *testing.T) {
	PreOrderTraversal(tree)
}

func TestMidOrderTraversal(t *testing.T) {
	MidOrderTraversal(tree)
}

func TestMidOrderTraversalIter(t *testing.T) {
	MidOrderTraversalIter(tree)
}

