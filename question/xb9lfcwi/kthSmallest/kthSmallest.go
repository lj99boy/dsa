package main

 type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

var seq = 0
var tK int
var res = 0

func main()  {
	root:=&TreeNode{Val: 5,Left: &TreeNode{Val: 3,Left: &TreeNode{Val: 2,Left: &TreeNode{Val: 1}},Right: &TreeNode{Val: 4}},Right: &TreeNode{Val: 6}}

	kthSmallest(root,3)
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return res
	}

	seq = 0
	res = 0
	tK = k

	midReverse(root)
	return res
}

func midReverse(root  *TreeNode)  {
	if root.Left != nil {
		midReverse(root.Left)
	}

	seq +=1

	if seq == tK {
		res = root.Val
	}
	if root.Right != nil {
		midReverse(root.Right)
	}
}