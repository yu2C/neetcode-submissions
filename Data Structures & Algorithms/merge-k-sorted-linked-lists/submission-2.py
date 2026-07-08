# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:    
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        if not lists:
            return
        
        while len(lists) > 1:
            merged = []
            for i in range(0, len(lists), 2):
                l = lists[i]
                r = None
                if i+1 < len(lists):
                    r = lists[i+1]

                merged.append(self.merge2Lists(l, r))
            lists = merged
        return lists[0]

    def merge2Lists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if not list1:
            return list2
        if not list2:
            return list1
        
        d = ListNode()
        cur = d

        while list1 and list2:
            if list1.val < list2.val:
                cur.next = list1
                list1 = list1.next
            else:
                cur.next = list2
                list2 = list2.next
            cur = cur.next
        
        if list1:
            cur.next = list1
        if list2:
            cur.next = list2
        
        return d.next