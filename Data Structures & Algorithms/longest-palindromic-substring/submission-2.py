class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) <= 1:
            return s

        self.start = 0
        self.max_len = 1

        for i in range(len(s)):
            # odd-length palindrome
            self.expand(s, i, i)

            # even-length palindrome
            self.expand(s, i, i + 1)

        return s[self.start:self.start + self.max_len]

    def expand(self, s: str, left: int, right: int) -> None:
        while left >= 0 and right < len(s) and s[left] == s[right]:
            length = right - left + 1

            if length > self.max_len:
                self.start = left
                self.max_len = length

            left -= 1
            right += 1