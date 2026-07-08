func minWindow(s string, t string) string {
   tMap := make(map[byte]int)
   for i := range(len(t)) {
        tMap[t[i]]++
   }

   l := 0
   haveCount := 0
   needCount := len(tMap)
   resLen := len(s) + 1
   res := []int{-1, -1}
   sMap := make(map[byte]int)

   for r := 0; r < len(s); r++ {
        sMap[s[r]]++

        if _, ok := tMap[s[r]]; ok && sMap[s[r]] == tMap[s[r]] {
            haveCount++
        }
        
        for haveCount == needCount {
            if r-l+1 < resLen {
                resLen = r-l+1
                res = []int{l, r}
            }

            sMap[s[l]]--

            if _, ok := tMap[s[l]]; ok && sMap[s[l]] < tMap[s[l]] {
                haveCount--
            }
            
            l++
            
        }
   }

    if resLen == len(s) + 1 {
        return ""
    }

   return s[res[0]:res[1]+1]
}
