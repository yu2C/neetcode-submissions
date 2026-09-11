func missingNumber(nums []int) int {
    xor := len(nums)

    for i, num := range nums {
        xor ^= i
        xor ^= num
    }

    return xor

}
