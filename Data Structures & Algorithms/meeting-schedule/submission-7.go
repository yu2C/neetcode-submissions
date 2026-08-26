/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
    if len(intervals) == 0 {
        return true
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    prevEnd := intervals[0].end
    for _, interval := range intervals[1:] {
        if prevEnd > interval.start {
            return false
        }
        prevEnd = interval.end
    }
    return true
}
