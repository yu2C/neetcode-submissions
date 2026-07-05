func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, 0)

    for _, v := range nums {
        count[v]++
    }

    type Pair struct {
        Num int
        Freq int
    }
    pairs := make([]Pair, 0)

    for key, value := range count {
        pairs = append(pairs, Pair{key, value})    
    }
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].Freq > pairs[j].Freq
    })

    res := make([]int, 0)

    for i := 0 ; i < k ; i++ {
        res = append(res, pairs[i].Num)
    }

    return res
}
