package numTrees


var resMem [][]int

func numTrees(n int) int {
	resMem = make([][]int,n+2)
	for i:=1;i<=n+1;i++{
		resMem[i] = make([]int,n+2)
	}
	return count(1,n)
}

func count(left int,right int)int  {

	if left > right {
		return 1
	}

	var res,leftRes,rightRes int
	for i:=left;i<=right;i++{
		if resMem[left][i-1]!=0 {
			leftRes = resMem[left][i-1]
		} else{
			leftRes = count(left,i-1)
			resMem[left][i-1]=leftRes
		}

		if resMem[i+1][right]!=0 {
			rightRes = count(i+1,right)
		} else{
			rightRes = count(i+1,right)
			resMem[i+1][right]=rightRes
		}

		res += leftRes * rightRes
	}

	return res
}