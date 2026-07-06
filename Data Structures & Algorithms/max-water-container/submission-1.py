class Solution:
	def maxArea(self, heights: List[int]) -> int:
		l, r = 0, len(heights)-1
		v = 0

		while l < r:
			w = r - l
			h = min(heights[l], heights[r])
			cur_v = w * h

			v = max(v, cur_v)

			if heights[l] < heights[r]:
				l += 1
			else:
				r -= 1
		return v