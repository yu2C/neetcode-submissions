func hasDuplicate(nums []int) bool {
    HashMap := make(map[int]bool)

	for _, num := range nums {
		if HashMap[num] {
			return true
		}

		HashMap[num] = true
	}
	return false

}
