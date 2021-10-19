package generateTrees

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }


func generateTrees(n int) []*TreeNode {
	if n <1 {
		return []*TreeNode{nil}
	}
	return build(1,n)
}

func build(left ,right int) []*TreeNode {
	rootTmpRes := make([]*TreeNode,0)

	if left>right {
		rootTmpRes = append(rootTmpRes,nil)
		return rootTmpRes
	} else {
		for i:=left;i<=right;i++{
			LeftRes := build(left,i-1)
			rightRes := build(i+1,right)
			for _,l:=range LeftRes{
				for _,r:=range rightRes{
					root := &TreeNode{}
					root.Val = i
					root.Left = l
					root.Right = r
					rootTmpRes = append(rootTmpRes,root)
				}
			}
		}
	}

	return rootTmpRes
}