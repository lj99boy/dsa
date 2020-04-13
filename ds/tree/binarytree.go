package tree

import (
	stack2 "dsa/ds/stack"
	"fmt"
)

type ChainBinaryTree struct {
	Val   interface{}
	Left  *ChainBinaryTree
	Right *ChainBinaryTree
}

type ArrayBinaryTree struct {
	Holder map[int]interface{}
}

func PreOrderTraversal(tree *ChainBinaryTree) {
	if tree != nil {
		fmt.Printf("%#v", tree.Val)
		PreOrderTraversal(tree.Left)
		PreOrderTraversal(tree.Right)
	}
}

func MidOrderTraversal(tree *ChainBinaryTree) {
	if tree != nil {
		MidOrderTraversal(tree.Left)
		fmt.Printf("%#v", tree.Val)
		MidOrderTraversal(tree.Right)
	}
}

func MidOrderTraversalIter(tree *ChainBinaryTree) {
	if tree == nil {
		return
	}
	stack := stack2.New()
	point := tree
	for ; !stack.IsEmpty() || point != nil; {
		for ; point != nil; {
			stack.Push(point)
			point = point.Left
		}

		point = stack.Pop().(*ChainBinaryTree)
		fmt.Printf("%v", point.Val)
		point = point.Right
	}
}
