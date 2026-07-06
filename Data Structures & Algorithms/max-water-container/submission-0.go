func maxArea(heights []int) int {
    l, r := 0, len(heights)-1

	v := 0

	for l < r {
		maxH := min(heights[l], heights[r])
		curV := (r-l) * maxH
		
		v = max(v, curV)
		
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return v
}