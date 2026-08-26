/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
 import (
	"slices"
 )

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	start := make([]int, 0, len(intervals))
	end := make([]int, 0, len(intervals))

	for i := range intervals {
		start = append(start, intervals[i].start)
		end = append(end, intervals[i].end)
	}

	slices.Sort(start)
	slices.Sort(end)

	s, e := 0, 0

	count := 0
	maxRoom := 0

	for s < len(intervals) {
		// next start < curEnd
		if start[s] < end[e] {
			count++
			maxRoom = max(maxRoom, count)
			s++
		} else { // next start >= curEnd
			count--
			e++
		}
	}	
	return maxRoom
}
