func countSubstrings(s string) int {
    var expand func(left int, right int) int
    expand = func(left int, right int) int {
        count := 0

        for left >= 0 && right < len(s) && s[left] == s[right] {
            count++
            left--
            right++
        }
        return count
    }
    res := 0
    for i := range s {
        // Odd palindrome
        res += expand(i, i)
        // Even palindrome
        res += expand(i, i+1)
    }
    return res
}
