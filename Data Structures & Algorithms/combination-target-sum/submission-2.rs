impl Solution {
    pub fn combination_sum(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        fn dfs(
            nums: &[i32],
            idx: usize,
            rem: i32,
            cur: &mut Vec<i32>,
            res: &mut Vec<Vec<i32>>
        ) {
            if rem == 0 {
                res.push(cur.clone());
                return;
            }

            if rem < 0 || nums.len() == idx {
                return;
            }

            cur.push(nums[idx]);
            dfs(nums, idx, rem - nums[idx], cur, res);
            cur.pop();

            dfs(nums, idx+1, rem, cur, res);

        }

        let mut res = Vec::new();
        let mut cur = Vec::new();

        dfs(
            &nums,
            0,
            target,
            &mut cur,
            &mut res
        );

        res
    }
}
