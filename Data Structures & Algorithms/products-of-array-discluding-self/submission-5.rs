impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut res = vec![1; n];

        for i in 1..n {
            res[i] = res[i-1] * nums[i-1];
        }

        let mut right = 1;

        for i in (0..n).rev() {
            res[i] *= right;
            right *= nums[i];
        }
        res
    }
}
