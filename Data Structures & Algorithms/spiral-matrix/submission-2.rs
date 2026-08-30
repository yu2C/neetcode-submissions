impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let rows = matrix.len();
        let cols = matrix[0].len();

        let mut res = Vec::with_capacity(rows * cols);

        let mut top: i32 = 0;
        let mut bottom: i32 = rows as i32 - 1;
        let mut left: i32 = 0;
        let mut right: i32 = cols as i32 - 1;

        while top <= bottom && left <= right {
            // top: left -> right
            for col in left..=right {
                res.push(matrix[top as usize][col as usize]);
            }
            top += 1;

            // right: top -> bottom
            for row in top..=bottom {
                res.push(matrix[row as usize][right as usize]);
            }
            right -= 1;

            if top <= bottom {
                // bottom: right -> left
                for col in (left..=right).rev() {
                    res.push(matrix[bottom as usize][col as usize]);
                }
                bottom -= 1;
            }

            if left <= right {
                // left: bottom -> top
                for row in (top..=bottom).rev() {
                    res.push(matrix[row as usize][left as usize])
                }
                left += 1;
            }
        }
        res
    }
}
