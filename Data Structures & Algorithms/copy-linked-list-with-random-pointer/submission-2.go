/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    cur := head

    for cur != nil {
        newNode := &Node{Val: cur.Val, Next: cur.Next}
        cur.Next = newNode
        cur = newNode.Next
    }

    cur = head 
    for cur != nil {
        if cur.Random != nil {
            cur.Next.Random = cur.Random.Next
        }
        cur = cur.Next.Next
    }

    cur = head
    newHead := head.Next
    for cur != nil {
        tmp := cur.Next
        cur.Next = tmp.Next
        if tmp.Next != nil {
            tmp.Next = tmp.Next.Next
        }
        cur = cur.Next
    }
    return newHead

}
