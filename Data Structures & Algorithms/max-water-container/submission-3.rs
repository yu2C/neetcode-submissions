impl Solution {
    pub fn max_area(heights: Vec<i32>) -> i32 {
		let mut l = 0usize;
		let mut r = heights.len()-1;

		let mut v:i32 = 0;

		while l < r {
			let h = heights[l].min(heights[r]);
			let w = r - l;
			let cur_v = h as i32 * w as i32;

			v = v.max(cur_v);

			if heights[l] < heights[r] {
				l += 1;
			} else {
				r -= 1;
			}
		}
		v
    }
}
