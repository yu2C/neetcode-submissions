func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    boxes := make([]map[byte]bool, 9)


    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            val := board[r][c]

            if val == '.' {
                continue
            }

            box := (r/3)*3 + c/3

            if rows[r][val] || cols[c][val] || boxes[box][val] {
                return false
            }

            rows[r][val] = true
            cols[c][val] = true
            boxes[box][val] = true
        }
    }

    return true
}
