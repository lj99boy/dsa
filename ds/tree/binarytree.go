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
		fmt.Printf("%#v ", tree.Val)
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

//先中后遍历路径都是一样的 自己 左 右， 先序是第一次遇见自己就处理，中序是第二次，后续是第三次
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

func PreOrderTraversalIter(tree *ChainBinaryTree) {
	if tree == nil {
		return
	}
	stack := stack2.New()
	stack.Push(tree)
	for ; !stack.IsEmpty(); {
		point := stack.Pop().(*ChainBinaryTree)
		if point == nil {
			continue
		}
		fmt.Printf("%v", point.Val)
		stack.Push(point.Right)
		stack.Push(point.Left)
	}
}

func PreOrderTraversalIter2(tree *ChainBinaryTree) {
	if tree == nil {
		return
	}
	stack := stack2.New()
	point := tree
	for ; !stack.IsEmpty() || point != nil; {
		for ; point != nil; {
			fmt.Printf("%v", point.Val)
			stack.Push(point)
			point = point.Left
		}

		point = stack.Pop().(*ChainBinaryTree)
		point = point.Right
	}
}
