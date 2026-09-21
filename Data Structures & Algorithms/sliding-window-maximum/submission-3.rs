impl Solution {
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let k = k as usize;

        let mut res = Vec::with_capacity(n - k + 1);
        let mut q: VecDeque<usize> = VecDeque::new();

        for r in 0..n {
            if let Some(&front) = q.front() {
                if front + k <= r {
                    q.pop_front();
                }
            }

            while let Some(&back) = q.back() {
                if nums[back] <= nums[r] {
                    q.pop_back();
                } else {
                    break;
                }
            }

            q.push_back(r);

            if r + 1 >= k {
                res.push(nums[*q.front().unwrap()]);
            }
        }

        res
    }
}
