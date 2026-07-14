type Trie struct {
    child [26]*Trie
    word string
}

func findWords(board [][]byte, words []string) []string {
    if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
        return []string{}
    }

    node := &Trie{}

    for i := 0; i < len(words); i++ {
        cur := node
        for j := 0; j < len(words[i]); j++ {
            index := int(words[i][j] - 'a')
            if cur.child[index] == nil {
                cur.child[index] = &Trie{}
            }

            cur = cur.child[index]
        }
        cur.word = words[i]
    }

    rows := len(board)
    cols := len(board[0])
    res := make([]string, 0)

    var dfs func(row int, col int, node *Trie)

    dfs = func(row int, col int, node *Trie) {
        if row < 0 || row >= rows || col < 0 || col >= cols || node == nil {
           return
        }

        ch := board[row][col]
        if ch == '#' {
            return
        }

        index := int(ch - 'a')
        next := node.child[index]
        if next == nil {
            return
        }

        if next.word != "" {
            res = append(res, next.word)
            next.word = ""
        }

        board[row][col] = '#'

        dfs(row-1, col, next)
        dfs(row+1, col, next)
        dfs(row, col-1, next)
        dfs(row, col+1, next)

        board[row][col] = ch
    }

    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            index := int(board[row][col] - 'a')
            if node.child[index] != nil {
                dfs(row, col, node)
            }
        }
    }

    return res
}
