func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	count := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j, row, col)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j, row, col int) {
	if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] != '1' {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '0'
		dfs(grid, i-1, j, row, col)
		dfs(grid, i+1, j, row, col)
		dfs(grid, i, j-1, row, col)
		dfs(grid, i, j+1, row, col)
	}
}