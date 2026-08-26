"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

class Solution:
	def minMeetingRooms(self, intervals: List[Interval]) -> int:
		starts, ends = [], []
		for interval in intervals:
			starts.append(interval.start)
			ends.append(interval.end)
		
		starts.sort()
		ends.sort()

		s, e, count, maxRoom = 0, 0, 0, 0

		while s < len(starts) and e < len(ends):
			if starts[s] < ends[e]:
				count += 1
				maxRoom = max(maxRoom, count)
				s += 1
			else:
				count -= 1
				e += 1
		return maxRoom