impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut count = HashMap::new();

        for n in nums {
            *count.entry(n).or_insert(0) += 1;
        }

        let mut pairs = Vec::new();

        for (key, value) in count {
            pairs.push((key, value));
        }

        pairs.sort_by(|a, b| b.1.cmp(&a.1));

        let mut res = Vec::new();

        for i in 0..k as usize {
            res.push(pairs[i].0);
        }

        res
    }
}
