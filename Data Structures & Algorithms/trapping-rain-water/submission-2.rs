impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
		let n = height.len();

		if n == 0 {
			return 0
		}

		let mut l_max = vec![0; n];
		let mut r_max = vec![0; n];

		l_max[0] = height[0];
		for i in (1..n) {
			l_max[i] = l_max[i-1].max(height[i]);
		}

		r_max[n-1] = height[n-1];
		for i in (0..n-1).rev() {
			r_max[i] = r_max[i+1].max(height[i]);
		}

		let mut water = 0;
		for i in (0..n) {
			water += l_max[i].min(r_max[i]) - height[i]
		}

		water
    }
}
