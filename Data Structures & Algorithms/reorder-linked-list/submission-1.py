# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        if not head or not head.next:
            return

        fast = slow = head

        while fast.next and fast.next.next:
            fast = fast.next.next
            slow = slow.next
        
        mid = slow.next
        slow.next = None

        cur = mid
        pre = None
        while cur:
            tmp = cur.next
            cur.next = pre

            pre = cur
            cur = tmp
        
        l, r = head, pre
        while r:
            next_l = l.next
            next_r = r.next

            l.next = r
            r.next = next_l

            l, r = next_l, next_r
