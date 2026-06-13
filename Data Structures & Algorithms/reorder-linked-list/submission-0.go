/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    fast, slow := head, head
    // Find Middle
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }

    mid := slow.Next
    slow.Next = nil

    var pre *ListNode
    cur := mid
    for cur != nil {
        tmp := cur.Next
        cur.Next = pre

        pre = cur
        cur = tmp
    }

    left, right := head, pre
    for right != nil {
        nextL := left.Next
        nextR := right.Next

        left.Next = right
        right.Next = nextL

        left = nextL
        right = nextR
    }

}
