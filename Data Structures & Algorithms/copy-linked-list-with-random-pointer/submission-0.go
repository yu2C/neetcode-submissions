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

    hash := make(map[*Node]*Node)

    cur := head
    for cur != nil {
        hash[cur] = &Node{Val: cur.Val}
        cur = cur.Next
    }

    cur = head
    for cur != nil {
        newNode := hash[cur]
        newNode.Next = hash[cur.Next]
        newNode.Random = hash[cur.Random]
        cur = cur.Next
    }

    return hash[head]
}
