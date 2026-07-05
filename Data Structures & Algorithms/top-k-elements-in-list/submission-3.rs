impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let mut counts: HashMap<i32, usize> = HashMap::new();

        for num in nums {
            *counts.entry(num).or_insert(0) += 1;
        }

        let mut bucket = vec![Vec::new(); n+1];

        for (num, freq) in counts {
            bucket[freq].push(num);
        }

        let mut res = Vec::new();

        for freq in (1..bucket.len()).rev() {
            for &num in &bucket[freq] {
                res.push(num);
                if res.len() == k as usize {
                    return res;
                }
            }

        }
        res
        
    }
}
