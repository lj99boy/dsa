package maxSumBST

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

 var max int

func maxSumBST(root *TreeNode) int {
	max = 0

	reverse(root)

	return max
}

func reverse(root *TreeNode) (isBst bool,minVal int,maxVal int,valSum int) {
	if root == nil {
		return true,4 * 10^4,-4 * 10^4,0
	}

	lIsBst,lMinVal,lMaxVal,lValSum := reverse(root.Left)
	rIsBst,rMinVal,rMaxVal,rValSum := reverse(root.Right)



	if lIsBst && rIsBst && root.Val>lMaxVal &&root.Val<rMinVal {
		curRootSum := lValSum +rValSum + root.Val

		if max < curRootSum {
			max = curRootSum
		}

		if root.Val < lMinVal {
			lMinVal = root.Val
		}

		if root.Val > rMaxVal {
			rMaxVal = root.Val
		}

		return true,lMinVal,rMaxVal,curRootSum
	} else {
		return false,0,0,0
	}
}