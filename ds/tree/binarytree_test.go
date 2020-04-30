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
			Left: &ChainBinaryTree{
				Val: 8,
				Left: &ChainBinaryTree{
					Val: 9,
				},
			},
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

func TestPreOrderTraversalIter(t *testing.T) {
	PreOrderTraversalIter(tree)
}

func TestPreOrderTraversalIter2(t *testing.T) {
	PreOrderTraversalIter2(tree)
}

func TestFindMaxDepth(t *testing.T) {
	println(FindMaxDepth(tree))
}


func TestFindMaxDepthByStack(t *testing.T) {
	println(FindMaxDepthByStack(tree))
}