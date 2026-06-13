type Node struct {
    key, val int
    prev, next *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head, tail *Node
}

func Constructor(capacity int) LRUCache {
    h, t := &Node{}, &Node{}
    h.next = t
    t.prev = h

    return LRUCache{
        capacity: capacity,
        cache: make(map[int]*Node),
        head: h,
        tail: t,
    }
}

func (this *LRUCache) Get(key int) int {
    cur, ok := this.cache[key]
    if !ok {
        return -1
    }
    this.moveToHead(cur)
    return cur.val
    
}

func (this *LRUCache) Put(key int, value int) {
    if cur, ok := this.cache[key]; ok {
        cur.val = value
        this.moveToHead(cur)
    } else {
        newNode := &Node{
            key: key,
            val: value,
        }
        this.cache[key] = newNode
        this.addToHead(newNode)

        if len(this.cache) > this.capacity {
            tar := this.tail.prev
            this.removeNode(tar)
            delete(this.cache, tar.key)
        }
    }
}

func (this *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *Node) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) moveToHead(node *Node) {
    this.removeNode(node)
    this.addToHead(node)
}

