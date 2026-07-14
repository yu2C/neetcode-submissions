type TrieNode struct {
    child [26]*TrieNode
    isWord bool
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
        root: &TrieNode{},
    }
}

func (this *WordDictionary) AddWord(word string)  {
    cur := this.root

    for i := 0; i < len(word); i++ {
        index := word[i] - 'a'

        if cur.child[index] == nil {
            cur.child[index] = &TrieNode{}
        }
        cur = cur.child[index]
    }

    cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
    var dfs func(node *TrieNode, index int) bool

    dfs = func(node *TrieNode, index int) bool {
        if node == nil {
            return false
        }

        if index == len(word) {
            return node.isWord
        }

        ch := word[index]

        if ch != '.' {
            chIndex := ch - 'a'
            return dfs(node.child[chIndex], index+1)
        }

        for _, c := range node.child {
            if c != nil && dfs(c, index+1) {
                return true
            }
        }
        
        return false
    }

    return dfs(this.root, 0)
}
