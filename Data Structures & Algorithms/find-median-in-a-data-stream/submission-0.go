type MedianFinder struct {
    nums []int
}


func Constructor() MedianFinder {
    return MedianFinder{
        nums: make([]int, 0),
    }
}


func (this *MedianFinder) AddNum(num int)  {
    this.nums = append(this.nums, num)
}


func (this *MedianFinder) FindMedian() float64 {
    sort.Slice(this.nums, func(i, j int) bool {
        return this.nums[i] < this.nums[j]
    })
    n := len(this.nums)

    if n%2 == 1 {
        return float64(this.nums[n/2])
    }

    left := this.nums[n/2-1]
    right := this.nums[n/2]
    return float64(left+right) / 2.0
}
