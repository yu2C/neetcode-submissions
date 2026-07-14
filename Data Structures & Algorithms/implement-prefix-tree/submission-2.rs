struct PrefixTree {
    child: [Option<Box<PrefixTree>>; 26],
    is_end: bool,
}

impl PrefixTree {
    pub fn new() -> Self {
        Self {
            child: std::array::from_fn(|_| None),
            is_end: false,
        }
    }

    pub fn insert(&mut self, word: String) {
        let mut cur = self;

        for byte in word.bytes() {
            let index = (byte - b'a') as usize;

            if cur.child[index].is_none() {
                cur.child[index] = Some(Box::new(PrefixTree::new()));
            }

            cur = cur.child[index].as_mut().unwrap();
        }

        cur.is_end = true;
    }

    pub fn search(&self, word: String) -> bool {
        let mut cur = self;

        for byte in word.bytes() {
            let index = (byte - b'a') as usize;

            match cur.child[index].as_ref() {
                Some(child) => {cur = child},
                None => return false,
            }
        }
        cur.is_end
    }

    pub fn starts_with(&self, prefix: String) -> bool {
        let mut cur = self;

        for byte in prefix.bytes() {
            let index = ( byte - b'a' ) as usize;

            match cur.child[index].as_ref() {
                Some(child) => {cur = child},
                None => return false,
            }
        }
        return true
    }
}
