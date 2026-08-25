func insert(intervals [][]int, newInterval []int) [][]int {
    res := make([][]int, 0, len(intervals)+1)

    for i, interval := range intervals {
        // interval completed before newInterval
        if newInterval[0] > interval[1] {
            res = append(res, interval)
            continue
        }

        // interval completed after newInterval
        if newInterval[1] < interval[0] {
            res = append(res, newInterval)
            res = append(res, intervals[i:]...)
            return res
        }

        // overlap
        newInterval[0] = min(newInterval[0], interval[0])
        newInterval[1] = max(newInterval[1], interval[1])
    }

    res = append(res, newInterval)

    return res
}
