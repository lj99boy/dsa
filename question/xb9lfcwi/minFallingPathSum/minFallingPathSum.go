package minFallingPathSum

import "math"

var dpMemo [][]int

func minFallingPathSum(matrix [][]int) int {
	dpMemo = make([][]int,len(matrix))
	for i:=0;i<len(matrix);i++{
		dpMemo[i] = make([]int,len(matrix[i]))
		for j:=0;j<len(matrix[i]);j++{
			dpMemo[i][j] = -2
		}
	}

	res := 99999
	for i:=0;i<len(matrix[0]);i++{
		tRes := dp(matrix,len(matrix)-1,i)
		if tRes<res {
			res =tRes
		}
	}

	return res
}

func dp(matrix [][]int,i,j int) int {
	if j<0||j>len(matrix[0])-1{
		return 99999
	}

	if i<1{
		return matrix[i][j]
	}

	if dpMemo[i][j] != -2 {
		return dpMemo[i][j]
	}

	topLeft := dp(matrix,i-1,j-1)
	top := dp(matrix,i-1,j)
	topRight := dp(matrix,i-1,j+1)

	nowVal := int(math.Min(math.Min(float64(topLeft), float64(top)), float64(topRight)))+matrix[i][j]

	dpMemo[i][j] = nowVal

	return nowVal
}