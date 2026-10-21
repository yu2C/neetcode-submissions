class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        lookup = {
            ')':'(',
            ']':'[',
            '}':'{',
        }

        for char in s:
            ## closing
            if char in lookup:
                if len(stack) == 0:
                    return False
                
                top = stack[-1]
                if top != lookup[char]:
                    return False
                
                stack.pop()
            else:
                stack.append(char)
        return len(stack) == 0

