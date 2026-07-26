func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)

    indegree := make([]int, numCourses)

    for _, value := range prerequisites {
        graph[value[0]] = append(graph[value[0]], value[1])
        indegree[value[1]]++
    }

    queue := make([]int, 0, numCourses)

    for course := 0; course < numCourses; course++ {
        if indegree[course] == 0 {
            queue = append(queue, course)
        }
    }

    completed := 0
    head := 0

    for head < len(queue) {
        course := queue[head]
        head++
        completed++

        for _, nextCourse := range graph[course] {
            indegree[nextCourse]--

            if indegree[nextCourse] == 0 {
                queue = append(queue, nextCourse)
            }
        }
    }

    return completed == numCourses


}
