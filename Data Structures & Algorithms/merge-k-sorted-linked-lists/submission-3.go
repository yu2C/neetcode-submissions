/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    h := &MinHeap{}
    heap.Init(h)

    for _, node := range(lists) {
        if node != nil {
            heap.Push(h, node)
        }
    }

    dummy := &ListNode{}
    cur := dummy

    for h.Len() > 0 {
        node := heap.Pop(h).(*ListNode)

        cur.Next = node
        cur = cur.Next

        if node.Next != nil {
            heap.Push(h, node.Next)
        }

    }
    return dummy.Next
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
    return len(h)
} 

func (h MinHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(i any) {
    *h = append(*h, i.(*ListNode))
}

func (h *MinHeap) Pop() any {
    old := *h
    res := old[len(old)-1]
    *h = old[:len(old)-1]
    return res
}