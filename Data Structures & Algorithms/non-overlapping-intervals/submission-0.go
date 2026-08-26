func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] > intervals[i][0] {
			count++
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
		}
	}
	return count
}
