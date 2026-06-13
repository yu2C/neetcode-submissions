/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var dummy *ListNode
	if head == nil {
		return dummy
	}
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = dummy

		dummy = cur
		cur = tmp
	}
	return dummy
}
