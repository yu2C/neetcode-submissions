func pacificAtlantic(heights [][]int) [][]int {
    rows := len(heights)
    if rows == 0 {
        return [][]int{}
    }
    cols := len(heights[0])
    if cols == 0 {
        return [][]int{}
    }

    pac := make([][]bool, rows)
    atl := make([][]bool, rows)

    for row := 0; row < rows; row++ {
        pac[row] = make([]bool, cols)
        atl[row] = make([]bool, cols)
    }

    directions := [][2]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }

    var dfs func(row, col int, visited [][]bool)
    dfs = func(row, col int, visited [][]bool) {
        visited[row][col] = true

        for _, direction := range directions {
            nextRow := row + direction[0]
            nextCol := col + direction[1]

            if nextRow < 0 || nextRow >= rows ||
                nextCol < 0 || nextCol >= cols {
                    continue
            }

            if visited[nextRow][nextCol] {
                continue
            }

            if heights[nextRow][nextCol] < heights[row][col] {
                continue
            }

            dfs(nextRow, nextCol, visited)
        }
    }

    for col := 0; col < cols; col++ {
        dfs(0, col, pac)
        dfs(rows-1, col, atl)
    }

    for row := 0; row < rows; row++ {
        dfs(row, 0, pac)
        dfs(row, cols-1, atl)
    }

    res := make([][]int, 0)

    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if pac[row][col] && atl[row][col] {
                res = append(res, []int{row, col})
            }
        }
    }

    return res
        
}
