func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
    for _, s := range(strs) {
        count := [26]int{}
        for _, char := range(s) {
            count[int(char - 'a')]++
        }
        m[count] = append(m[count], s)
    }
    res := make([][]string, 0)
    for _, v := range m {
        res = append(res, v)
    }
    return res
}