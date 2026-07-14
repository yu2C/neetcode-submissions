struct Trie {
    child: [Option<Box<Trie>>; 26],
    is_word: bool,
}

impl Trie {
    fn new() -> Self {
        Self {
            child: std::array::from_fn(|_| None),
            is_word: false,
        }
    }
}

struct WordDictionary {
    root: Trie,
}

impl WordDictionary {
    pub fn new() -> Self {
        Self {
            root: Trie::new(),
        }    
    }

    pub fn add_word(&mut self, word: String) {
        let mut cur = &mut self.root;

        for byte in word.bytes() {
            let index = ( byte - b'a' ) as usize;

            cur = cur.child[index]
                .get_or_insert_with(|| Box::new(Trie::new()));
        }
        cur.is_word = true;
    }

    pub fn search(&self, word: String) -> bool {
        fn dfs(word: &[u8], node: &Trie, index: usize) -> bool {
            if index == word.len() {
                return node.is_word;
            }

            let byte = word[index];

            if byte != b'.' {
                let child_index = (byte - b'a') as usize;

                return node.child[child_index]
                    .as_ref()
                    .is_some_and(|child| dfs(word, child, index+1));
            }

            node.child
                .iter()
                .flatten()
                .any(|child| dfs(word, child.as_ref(), index+1))
        }
        dfs(word.as_bytes(), &self.root, 0)
    }
}
