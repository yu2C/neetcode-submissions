func foreignDictionary(words []string) string {
    adjList := make(map[byte]map[byte]bool)
    indegree := make(map[byte]int)

    for _, word := range words {
        for i := 0; i < len(word); i++ {
            ch := word[i]

            if _, exists := adjList[ch]; !exists {
                adjList[ch] = make(map[byte]bool)
            }

            if _, exists := indegree[ch]; !exists {
                indegree[ch] = 0
            }

        } 
    }
    
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]

        length := min(len(word1), len(word2))
        foundDiff := false

        for j := 0; j < length; j++ {
            if word1[j] == word2[j] {
                continue
            }

            from := word1[j]
            to := word2[j]

            if !adjList[from][to] {
               adjList[from][to] = true
               indegree[to]++ 
            }

            foundDiff = true
            break
        }

        if !foundDiff && len(word1) > len(word2) {
            return ""
        }
    }

    // node -> queue
    queue := make([]byte, 0)

    for ch, degree := range indegree {
        if degree == 0 {
            queue = append(queue, ch)
        }
    }

    // topological sort
    res := make([]byte, 0, len(indegree))
    head := 0

    for head < len(queue) {
        node := queue[head]
        head++

        res = append(res, node)

        for neighbor := range adjList[node] {
            indegree[neighbor]--

            if indegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    if len(res) != len(indegree) {
        return ""
    }

    return string(res)
}
