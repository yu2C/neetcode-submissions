impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let set: HashSet<i32> = nums.into_iter().collect();
        let mut long = 0;

        for &n in &set {
            let mut l = 1;

            while set.contains(&(n+l)) {
                l += 1
            }

            long = long.max(l);
        }
        long
    }
}
