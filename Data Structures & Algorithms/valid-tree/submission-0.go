func validTree(n int, edges [][]int) bool {
    if n <= 0 {
        return false
    }
    if len(edges) != n - 1 {
        return false
    }

    graph := make([][]int, n)

    for _, edge := range edges {
        if len(edge) != 2 {
            return false
        }

        node1 := edge[0]
        node2 := edge[1]

        if node1 < 0 || node1 >= n || node2 < 0 || node2 >= n {
            return false
        }

        graph[node1] = append(graph[node1], node2)
        graph[node2] = append(graph[node2], node1)
    }

    visited := make([]bool, n)

    var dfs func(node int)
    dfs = func(node int) {
        visited[node] = true

        for _, neighbor := range graph[node] {
            if visited[neighbor] {
                continue
            }

            dfs(neighbor)
        }
    }

    dfs(0)

    for node := 0; node < n; node++ {
        if !visited[node] {
            return false
        }
    }

    return true


}
