"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        if node is None:
            return None
        
        cloned: dict[Node, Node] = {
            node: Node(node.val)
        }

        queue = deque([node])

        while queue:
            cur = queue.popleft()

            for neighbor in cur.neighbors:
                if neighbor not in cloned:
                    cloned[neighbor] = Node(neighbor.val)
                    queue.append(neighbor)

                cloned[cur].neighbors.append(cloned[neighbor])

        return cloned[node]