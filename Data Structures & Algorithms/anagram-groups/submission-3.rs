impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut m: HashMap<
            [usize; 26], 
            Vec<String>
            > = HashMap::new();
        

        for s in strs {
            let mut count = [0usize; 26];

            for ch in s.chars() {
                let idx = ch as usize - 'a' as usize;
                count[idx] += 1;
            }

            m.entry(count)
                .or_insert(Vec::new())
                .push(s);

        }

        m.into_values().collect()
    }
}
