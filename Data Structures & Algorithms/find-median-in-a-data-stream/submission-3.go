type MedianFinder struct {
    small MaxHeap
    large MinHeap
}

type MinHeap []int

func (h MinHeap) Len() int {
    return len(h)
}

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)

    value := old[n-1]
    *h = old[:n-1]

    return value
}

type MaxHeap []int

func (h MaxHeap) Len() int {
    return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
    old := *h
    n := len(old)

    value := old[n-1]
    *h = old[:n-1]

    return value
}

func Constructor() MedianFinder {
    small := MaxHeap{}
    large := MinHeap{}

    heap.Init(&small)
    heap.Init(&large)

    return MedianFinder{
        small: small,
        large: large,
    }
}


func (this *MedianFinder) AddNum(num int)  {
    if this.small.Len() == 0 || this.small[0] >= num {
        heap.Push(&this.small, num)
    } else {
        heap.Push(&this.large, num)
    }

    if this.small.Len() > this.large.Len() + 1 {
        val := heap.Pop(&this.small).(int)
        heap.Push(&this.large, val)
    } else if this.large.Len() > this.small.Len() {
        val := heap.Pop(&this.large).(int)
        heap.Push(&this.small, val)
    }
    
}


func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len() == this.large.Len() + 1 {
        return float64(this.small[0])
    }
    return float64(this.small[0] + this.large[0]) / 2.0
}
