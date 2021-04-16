package main

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfsGrid(grid, i, j)
				res++
			}
		}
	}
	return res
}

func dfsGrid(grid [][]byte, row int, col int) {
	if row >= len(grid) || col >= len(grid[0]) || row < 0 || col < 0 {
		return
	}
	if grid[row][col] != '1' {
		return
	}
	grid[row][col] = '2'
	dfsGrid(grid, row-1, col)
	dfsGrid(grid, row+1, col)
	dfsGrid(grid, row, col-1)
	dfsGrid(grid, row, col+1)
}
