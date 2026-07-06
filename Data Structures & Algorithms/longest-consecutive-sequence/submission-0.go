func longestConsecutive(nums []int) int {
    set := make(map[int]bool)

    for _, n := range nums {
        set[n] = true
    }

    long := 0

    for num := range set {
        if !set[num-1] {
            length := 1

            for set[num+length] {
                length++
            }

            if length > long {
                long = length
            }
        }
    }
    return long
}
