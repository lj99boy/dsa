package solve_n_queen

func SolveNQueens(n int) [][][]string {
	var res [][][]string

	var closure func([][]string,int)

	closure = func(tClassBoard [][]string,rowHeight int) {
		if rowHeight == n+1 {
			res = append(res, tClassBoard)
			return
		}

		for col := 0; col < n; col++ {
			var tmpClassBoard = tClassBoard
			if isVaild(tmpClassBoard,rowHeight,col){
				tmpClassBoard[rowHeight][col] = "Q"
				closure(tmpClassBoard,rowHeight+1)
				tmpClassBoard[rowHeight][col] = "."
			}
		}
	}

	return res
}

func isVaild(tClassBoard [][]string,row int,col int) bool {
	return true
}
