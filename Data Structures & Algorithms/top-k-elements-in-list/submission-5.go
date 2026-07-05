func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int, 0)

    for _, num := range nums {
        count[num]++
    }

    h := &MinHeap{}
    heap.Init(h)

    for num, freq := range count {
        heap.Push(h, Pair{num, freq})

        if h.Len() > k {
            heap.Pop(h)
        }
    }

    res := make([]int, 0)

    for h.Len() > 0 {
        item := heap.Pop(h).(Pair)
        res = append(res, item.Num)
    }
    
    return res
}

type Pair struct {
    Num int
    Freq int
}

type MinHeap []Pair

func (h *MinHeap) Len() int {
    return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
    return (*h)[i].Freq < (*h)[j].Freq
}

func (h *MinHeap) Swap(i, j int) {
    tmp := (*h)[i]
    (*h)[i] = (*h)[j]
    (*h)[j] = tmp
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    item := old[n-1]
    *h = old[:n-1]
    return item
}