type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var b strings.Builder

    for _, s := range strs {
        b.WriteString(strconv.Itoa(len(s)))
        b.WriteByte('#')
        b.WriteString(s)
    }
    return b.String()
}

func (s *Solution) Decode(encoded string) []string {
    res := []string{}

    i := 0

    for i < len(encoded) {
        j := i

        for encoded[j] != '#' {
            j++
        }

        length, _ := strconv.Atoi(encoded[i:j])

        start := j + 1
        end := start + length

        res = append(res, encoded[start:end])

        i = end
    }
    return res
}
