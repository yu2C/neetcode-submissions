func hammingWeight(n int) int {
	count := 0

	for i := 0; i < 32; i++ {
		mask := 1 << i

		if n & mask != 0 {
			count++
		}
	} 
	return  count

}
