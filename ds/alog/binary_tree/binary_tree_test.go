package binary_tree

import "testing"

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}


	println(BuildTree(preorder,inorder))
}
