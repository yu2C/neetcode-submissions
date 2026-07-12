impl Solution {
    pub fn exist(mut board: Vec<Vec<char>>, word: String) -> bool {
        let rows = board.len();
        let cols = board[0].len();
        let words: Vec<char> = word.chars().collect();

        fn dfs(
            board: &mut Vec<Vec<char>>, 
            word: &[char],
            row: i32, 
            col: i32, 
            idx: usize,
        ) -> bool {
                if idx == word.len() {
                    return true;
                }

                if row < 0
                    || row >= board.len() as i32
                    || col < 0
                    || col >= board[0].len() as i32
                    || board[row as usize][col as usize] != word[idx] {
                    return false;
                }

                let ori = board[row as usize][col as usize];
                board[row as usize][col as usize] = '#';

                let found = dfs(board, word, row-1, col, idx+1)
                    || dfs(board, word, row+1, col, idx+1)
                    || dfs(board, word, row, col-1, idx+1)
                    || dfs(board, word, row, col+1, idx+1);

                board[row as usize][col as usize] = ori;

                found
        }

        for row in 0..rows {
            for col in 0..cols {
                if dfs(
                    &mut board, &words, row as i32, col as i32, 0
                ) {
                    return true;
                }
            }
        }
        false
    }
}
