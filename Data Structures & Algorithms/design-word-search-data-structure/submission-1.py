class TrieNode:
    def __init__(self):
        self.child = {}
        self.is_word = False
class WordDictionary:

    def __init__(self):
        self.root = TrieNode()

    def addWord(self, word: str) -> None:
        cur = self.root
        for ch in word:
            if ch not in cur.child:
                cur.child[ch] = TrieNode()
            cur = cur.child[ch]
        cur.is_word = True

    def search(self, word: str) -> bool:
        def dfs(node: TrieNode, index: int) -> bool:
            if not node:
                return False
            if index == len(word):
                return node.is_word
            
            if word[index] != '.':
                return dfs(node.child.get(word[index]), index+1)
            
            for ch in node.child.values():
                if dfs(ch, index+1):
                    return True
            return False
                    
        return dfs(self.root, 0)
