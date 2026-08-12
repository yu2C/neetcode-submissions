func countComponents(n int, edges [][]int) int {
    graph := make([][]int, n)

    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }

    visited := make([]bool, n)

    var dfs func(int)
    dfs = func(node int) {
        visited[node] = true

        for _, neighbor := range graph[node] {
            if visited[neighbor] {
                continue
            }

            dfs(neighbor)
        }
    }

    count := 0

    for node := 0; node < n; node++ {
        if visited[node] {
            continue
        }
        count++
        dfs(node)
    }

    return count
}
