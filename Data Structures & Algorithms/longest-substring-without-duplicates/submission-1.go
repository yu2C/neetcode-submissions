func lengthOfLongestSubstring(s string) int {
    l := 0
    m := make(map[byte]bool)
    res := 0

    for r := 0 ; r < len(s) ; r++ {
        for m[s[r]] {
            delete(m, s[l])
            l++
        }
        m[s[r]] = true

        res = max(res, r-l+1)
    }
    return res

}
