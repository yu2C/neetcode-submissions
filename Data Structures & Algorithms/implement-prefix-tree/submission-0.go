type PrefixTree struct {
    children [26]*PrefixTree
    isEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
    cur := this

    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'
        if cur.children[index] == nil {
            cur.children[index] = &PrefixTree{}
        }
        cur = cur.children[index]
    }

    cur.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
    cur := this

    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'
        if cur.children[index] == nil {
            return false
        }
        cur = cur.children[index]
    }
    return cur.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
    cur := this

    for i := 0; i < len(prefix); i++ {
        index := prefix[i] - 'a'
        if cur.children[index] == nil {
            return false
        }
        cur = cur.children[index]
    }
    return true
}
