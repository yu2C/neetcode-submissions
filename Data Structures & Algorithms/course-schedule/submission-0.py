class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph: list[list[int]] = [[] for _ in range(numCourses)]

        indegree: list[int] = [0] * numCourses

        for course, prerequisite in prerequisites:
            graph[prerequisite].append(course)
            indegree[course] += 1
        
        queue: deque[int] = deque()

        for course in range(numCourses):
            if indegree[course] == 0:
                queue.append(course)

        completed = 0

        while queue:
            course = queue.popleft()
            completed += 1

            for next_course in graph[course]:
                indegree[next_course] -= 1

                if indegree[next_course] == 0:
                    queue.append(next_course)
        
        return completed == numCourses