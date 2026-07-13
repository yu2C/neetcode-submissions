class PrefixTree:

    def __init__(self):
        self.child = {}
        self.isEnd = False

    def insert(self, word: str) -> None:
        cur = self

        for ch in word:
            if ch not in cur.child:
                cur.child[ch] = PrefixTree()
            cur = cur.child[ch]
        cur.isEnd = True

    def search(self, word: str) -> bool:
        cur = self

        for ch in word:
            if ch not in cur.child:
                return False
            cur = cur.child[ch]
        return cur.isEnd

    def startsWith(self, prefix: str) -> bool:
        cur = self

        for ch in prefix:
            if ch not in cur.child:
                return False
            cur = cur.child[ch]
        return True