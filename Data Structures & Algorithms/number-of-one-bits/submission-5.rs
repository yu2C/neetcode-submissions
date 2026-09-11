impl Solution {
    pub fn hamming_weight(n: u32) -> i32 {
		let mut count = 0;

		for i in 0..32 {
			if n & (1u32<<i) != 0 {
				count += 1;
			}
		}

		count
    }
}
