impl Solution {
    pub fn encode(strs: Vec<String>) -> String {
        let mut res = String::new();

        for s in strs {
            res.push_str(&s.len().to_string());
            res.push('#');
            res.push_str(&s);
        }
        res
    }

    pub fn decode(s: String) -> Vec<String> {
        let bytes = s.as_bytes();
        let mut res = Vec::new();
        let mut i = 0;

        while i < bytes.len() {
            let mut j = i;

            while bytes[j] != b'#' {
                j += 1;
            }

            let length: usize =  s[i..j].parse().unwrap();

            let start = j + 1;
            let end = start + length;

            res.push(s[start..end].to_string());

            i = end;
        }
        res
    }
}
