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

func FindMaxDepthByStack(tree *ChainBinaryTree) int {
	if tree == nil {
		return 0
	}

	treeDepth := make(map[*ChainBinaryTree]int, 0)
	var stack []*ChainBinaryTree
	stack = append(stack, tree)
	maxDepth := 1
	tmpDepth := 1

	for ; len(stack) > 0; {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if val, ok := treeDepth[node]; ok {
			tmpDepth = val
		}
		if tmpDepth > maxDepth {
			maxDepth = tmpDepth
		}
		if node.Left != nil {
			treeDepth[node.Left] = tmpDepth + 1
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			treeDepth[node.Right] = tmpDepth + 1
			stack = append(stack, node.Right)
		}
	}
	return maxDepth
}

func FindMaxDepth(tree *ChainBinaryTree) int {
	if tree == nil {
		return 0
	}
	if tree.Left == nil && tree.Right == nil {
		return 1
	}
	return compareMax(FindMaxDepth(tree.Right), FindMaxDepth(tree.Left)) + 1
}

func compareMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
