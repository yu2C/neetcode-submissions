class Solution:
	def trap(self, height: List[int]) -> int:
		n = len(height)
		if n == 0:
			return 0
		
		l_max = [0] * n
		r_max = [0] * n

		l_max[0] = height[0]
		for i in range(1, n):
			l_max[i] = max(l_max[i-1], height[i])
		
		r_max[n-1] = height[n-1]
		for i in range(n-2, -1, -1):
			r_max[i] = max(r_max[i+1], height[i])

		water = 0

		for i in range(n):
			water += min(l_max[i], r_max[i]) - height[i]
		
		return water