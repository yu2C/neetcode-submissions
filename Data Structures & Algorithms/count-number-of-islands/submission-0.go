func numIslands(grid [][]byte) int {
    rows := len(grid)
    cols := len(grid[0])

    var dfs func(row, col int)  
    dfs = func(row, col int) {
        // edge
        if row < 0 || col < 0 ||
            row >= rows || col >= cols ||
            grid[row][col] != '1' {
            return
        }

        grid[row][col] = '0'

        dfs(row-1, col)
        dfs(row+1, col)
        dfs(row, col-1)
        dfs(row, col+1)
    }

    res := 0

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' {
                res++
                dfs(i, j)
            }
        }
    }
    return res
    
}
