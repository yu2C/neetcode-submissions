/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    head := dummy

    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            dummy.Next = list1
            list1 = list1.Next
        } else {
            dummy.Next = list2
            list2 = list2.Next
        }
        dummy = dummy.Next
    }
    dummy.Next = list1
    if list1 == nil {
        dummy.Next = list2
    }
    return head.Next
}