impl Solution {
    pub fn num_islands(mut grid: Vec<Vec<char>>) -> i32 {
        if grid.is_empty() || grid[0].is_empty() {
            return 0;
        }

        let rows = grid.len() as isize;
        let cols = grid[0].len() as isize;

        fn dfs(
            grid: &mut Vec<Vec<char>>,
            rows: isize,
            cols: isize,
            row: isize,
            col: isize,
        ) {
            if row < 0 || row >= rows || col < 0 || col >= cols {
                return;
            }

            let row_i = row as usize;
            let col_i = col as usize;

            if grid[row_i][col_i] != '1' {
                return;
            }

            grid[row_i][col_i] = '0';

            dfs(grid, rows, cols, row-1, col);
            dfs(grid, rows, cols, row+1, col);
            dfs(grid, rows, cols, row, col-1);
            dfs(grid, rows, cols, row, col+1);
        }

        let mut res = 0i32;

        for i in 0..rows {
            for j in 0..cols {
                if grid[i as usize][j as usize] == '1' {
                    res += 1;
                    dfs(&mut grid, rows, cols, i as isize, j as isize)
                }
            }
        }
        res
    }
}
