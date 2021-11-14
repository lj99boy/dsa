package bfs_minDepth

//https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/


  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func minDepth(root *TreeNode) int {

	if root==nil {
		return 0
	}

	treeHolder := make([]*TreeNode,0)
	treeHolder = append(treeHolder,root)

	depth:=0
	for ;len(treeHolder)>0;{
		depth++
		layerLen :=len(treeHolder)
		for i:=0;i<layerLen;i++{
			t:=treeHolder[i]
			if t.Left!=nil {
				treeHolder = append(treeHolder,t.Left)
			}
			if t.Right!=nil {
				treeHolder = append(treeHolder,t.Right)
			}

			if t.Left == nil && t.Right==nil {
				return depth
			}
		}
		treeHolder = treeHolder[layerLen:]
	}

	return depth
}