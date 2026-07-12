func exist(board [][]byte, word string) bool {
    rows := len(board)
    cols := len(board[0])

    var dfs func(row, col, idx int) bool

    dfs = func(row, col, idx int) bool {
        if idx == len(word) {
            return true
        }

        if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[idx] {
            return false
        }

        ori := board[row][col]
        board[row][col] = '#'

        found := dfs(row-1, col, idx+1) || dfs(row+1, col, idx+1) || dfs(row, col-1, idx+1) || dfs(row, col+1, idx+1)

        board[row][col] = ori

        return found
    }

    for i := range(rows) {
        for j := range(cols) {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}
