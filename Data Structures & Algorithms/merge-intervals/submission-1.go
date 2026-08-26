func merge(intervals [][]int) [][]int {
    res := make([][]int, 0, len(intervals))
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res = append(res, intervals[0])

    for i := 1; i < len(intervals); i++ {
        if res[len(res)-1][1] >= intervals[i][0] {
            res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
            continue
        }
        res = append(res, intervals[i])
    }
    return res
}
