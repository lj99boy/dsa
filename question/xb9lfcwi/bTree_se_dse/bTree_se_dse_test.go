package bTree_se_dse

import (
	"fmt"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	tree := &TreeNode{Val: 1,Left: &TreeNode{Val: 2},Right: &TreeNode{Val: 3,Left: &TreeNode{Val: 4},Right: &TreeNode{Val: 5}}}
	code:=Constructor()
	xx := code.serialize(tree)
	fmt.Println(xx)
	fmt.Println(code.deserialize(xx))

}
