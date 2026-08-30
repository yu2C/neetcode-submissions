impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();

        matrix.reverse();

        for i in 0..n {
            for j in (i+1)..n {
                let (top, bottom) = matrix.split_at_mut(j);

                std::mem::swap(
                    &mut top[i][j],
                    &mut bottom[0][i],
                );
            }
        }
    }
}
