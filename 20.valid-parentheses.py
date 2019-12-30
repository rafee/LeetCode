#
# @lc app=leetcode id=20 lang=python3
#
# [20] Valid Parentheses
#
# https://leetcode.com/problems/valid-parentheses/description/
#
# algorithms
# Easy (37.72%)
# Likes:    3803
# Dislikes: 190
# Total Accepted:    792.8K
# Total Submissions: 2.1M
# Testcase Example:  '"()"'
#
# Given a string containing just the characters '(', ')', '{', '}', '[' and
# ']', determine if the input string is valid.
#
# An input string is valid if:
#
#
# Open brackets must be closed by the same type of brackets.
# Open brackets must be closed in the correct order.
#
#
# Note that an empty string is also considered valid.
#
# Example 1:
#
#
# Input: "()"
# Output: true
#
#
# Example 2:
#
#
# Input: "()[]{}"
# Output: true
#
#
# Example 3:
#
#
# Input: "(]"
# Output: false
#
#
# Example 4:
#
#
# Input: "([)]"
# Output: false
#
#
# Example 5:
#
#
# Input: "{[]}"
# Output: true
#
#
#

# @lc code=start


class Solution:
    def isValid(self, s: str) -> bool:
        bracketStack = []
        bracketMap = {'(': ')', '{': '}', '[': ']'}
        for currentChar in s:
            if currentChar in bracketMap:
                bracketStack.append(bracketMap[currentChar])

            else:
                try:
                    if currentChar == bracketStack.pop():
                        continue
                    else:
                        return False
                except IndexError:
                    return False

        return bracketStack == []
        # @lc code=end
