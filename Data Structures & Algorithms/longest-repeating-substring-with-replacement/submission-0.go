func characterReplacement(s string, k int) int {
    l := 0
    maxFreq := 0

    count := make(map[byte]int)

    res := 0

    for r := 0 ; r < len(s) ; r++ {
        count[s[r]]++

        if count[s[r]] > maxFreq {
            maxFreq = count[s[r]]
        }

        for (r-l+1) - maxFreq > k {
            count[s[l]]--
            l++
        }

        if r-l+1 > res {
            res = r-l+1
        }
        
    }
    return res
}
