//寻找岛屿数量，可以用广度搜索，这里是取巧做法
package findisland

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	count := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] != '0' {
				count++
				mark(grid, r, c, row, col)
			}
		}
	}
	return count
}

func mark(grid [][]byte, r int, c int, row int, col int) {
	if r < 0 || c < 0 || r >= row || c >= col || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	mark(grid, r+1, c, row, col)
	mark(grid, r-1, c, row, col)
	mark(grid, r, c+1, row, col)
	mark(grid, r, c-1, row, col)
}
