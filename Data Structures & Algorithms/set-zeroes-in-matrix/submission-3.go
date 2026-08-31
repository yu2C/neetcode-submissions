func setZeroes(matrix [][]int) {
    firstRowZero := false
    firstColZero := false

    for col := 0; col < len(matrix[0]); col++ {
        if matrix[0][col] == 0 {
            firstRowZero = true
        }
    }

    for row := 0; row < len(matrix); row++ {
        if matrix[row][0] == 0 {
            firstColZero = true
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if firstColZero {
        for row := 0; row < len(matrix); row++ {
            matrix[row][0] = 0
        }
    }

    if firstRowZero {
        for col := 0; col < len(matrix[0]); col++ {
            matrix[0][col] = 0
        }
    }


}
