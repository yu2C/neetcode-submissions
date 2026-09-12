func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var count1  [26]int
    var count2  [26]int

    for i := 0; i < len(s1); i++ {
        count1[s1[i] - 'a']++
        count2[s2[i] - 'a']++
    }

    if count1 == count2 {
        return true
    }

    l := 0

    for r := len(s1); r < len(s2); r++ {
        count2[s2[r] - 'a']++
        count2[s2[l] - 'a']--

        l++

        if count1 == count2 {
            return true
        }
    }    
    return false
}
