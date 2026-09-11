impl Solution {
    pub fn count_bits(mut n: i32) -> Vec<i32> {
        let mut res = Vec::new();

        while n >= 0 {
            let mut count = 0;

            for i in 0..32 {
                if n & (1 << i) != 0 {
                    count += 1;
                }
            }

            res.push(count);

            n -= 1;
        }

        res.reverse();
        res
    }
}
