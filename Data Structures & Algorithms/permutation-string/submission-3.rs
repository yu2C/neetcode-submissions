impl Solution {
    pub fn check_inclusion(s1: String, s2: String) -> bool {
        if s1.len() > s2.len() {
            return false
        }

        let s1 = s1.as_bytes();
        let s2 = s2.as_bytes();

        let mut count1 = [0i32; 26];
        let mut count2 = [0i32; 26];

        for i in 0..s1.len() {
            count1[ ( s1[i] - b'a' ) as usize ] += 1;
            count2[ ( s2[i] - b'a' ) as usize ] += 1;
        }

        if count1 == count2 {
            return true;
        }

        let mut l = 0;

        for r in (s1.len()..s2.len()) {
            count2[ ( s2[r] - b'a' ) as usize ] += 1;
            count2[ ( s2[l] - b'a' ) as usize ] -= 1;

            l += 1;

            if count1 == count2 {
                return true
            }
        }

        false
    }
}
