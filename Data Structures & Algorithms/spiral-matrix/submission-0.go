func spiralOrder(matrix [][]int) []int {
    rows := len(matrix)
    cols := len(matrix[0])

    res := make([]int, 0, rows*cols)

    top := 0
    bottom := rows - 1
    left := 0
    right := cols - 1

    for top <= bottom && left <= right {

        // 1. top: left -> right
        for col := left; col <= right; col++ {
            res = append(res, matrix[top][col])
        }
        top++

        // 2. right: top -> bottom
        for row := top; row <= bottom; row++ {
            res = append(res, matrix[row][right])
        }
        right--

        if top <= bottom {
            // 3. bottom: right -> left
            for col := right; col >= left; col-- {
                res = append(res, matrix[bottom][col])
            }
            bottom--
        }

        if left <= right {
            // 4. left: bottom -> top
            for row := bottom; row >= top; row-- {
                res = append(res, matrix[row][left])
            }
            left++
        }
    }

    return res


}
