package tree

import (
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	vals := []int{2, 3, 7, 9, 18, 25}
	res := BuildHuffmanTree(vals)
	PreOrderTraversal(res)
}
