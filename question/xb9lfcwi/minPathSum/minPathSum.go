package minPathSum

import "math"

// https://leetcode-cn.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	if len(grid)<1{
		return 0
	}

	dpMemo = make([][]int,len(grid))

	for i:=0;i<len(grid);i++{
		dpMemo[i] = make([]int,len(grid[0]))
		for j:=0;j<len(grid[0]);j++{
			dpMemo[i][j] = 99998
		}
	}

	return dp(grid,len(grid)-1,len(grid[0])-1)
}

var dpMemo [][]int

func dp(grid[][]int ,i,j int) int {
	var res int

	if i-1<0 && j-1<0 {
		res= grid[i][j]
		dpMemo[i][j] = res
	}

	if dpMemo[i][j] != 99998 {
		return dpMemo[i][j]
	}

	if i-1<0{
		res = dp(grid,i,j-1)+grid[i][j]
		dpMemo[i][j] = res
	}

	if j-1<0{
		res = dp(grid,i-1,j)+grid[i][j]
		dpMemo[i][j] = res
	}

	if dpMemo[i][j] != 99998 {
		return dpMemo[i][j]
	}

	res = int(math.Min(float64(dp(grid,i-1,j) + grid[i][j]), float64(dp(grid,i,j-1)+grid[i][j])))
	dpMemo[i][j] = res
	return res
}


movl    %eax, -8(%rbp)
movl    -12(%rbp), %eax
movl    %eax, -4(%rbp)
movl    $0, %eax

int k = i + j;
int sum = add(11, 22,33, 44, 55, 66, 77, 88);
int m = k; // 为了观察 %rax Caller Save 寄存器的恢复