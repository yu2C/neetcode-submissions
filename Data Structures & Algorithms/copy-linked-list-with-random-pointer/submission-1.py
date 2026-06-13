"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None
        
        cur = head
        hash = {}
        while cur:
            hash[cur] = Node(cur.val)
            cur = cur.next
        
        cur = head
        while cur:
            newNode = hash[cur]
            newNode.next = hash.get(cur.next)
            newNode.random = hash.get(cur.random)
            cur = cur.next

        return hash[head]
