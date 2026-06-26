func twoSum(nums []int, target int) []int {
    seen := make(map[int]int, 0)

    for i, num := range nums {
        com := target - num

        if _, exist := seen[com]; exist {
            return []int{seen[com], i}
        }
        seen[num] = i
    }
    return nil
}
