impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;

        let mut xor = 0;

        for i in 1..=n {
            xor ^= i
        }

        for num in nums {
            xor ^= num;
        }
        
        xor
    }
}
