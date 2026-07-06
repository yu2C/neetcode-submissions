func productExceptSelf(nums []int) []int {
    l := make([]int, len(nums))
    r := make([]int, len(nums))
    res := make([]int, len(nums))

    l[0] = 1
    r[len(nums)-1] = 1

    for i := 1 ; i < len(nums) ; i++ {
        l[i] = l[i-1] * nums[i-1]
    
    }
    res[len(nums)-1] = l[len(nums)-1] * r[len(nums)-1]
    for i := len(nums)-2 ; i >= 0 ; i-- {
        r[i] = r[i+1] * nums[i+1]
        res[i] = r[i] * l[i]
    }


    return res
} 
