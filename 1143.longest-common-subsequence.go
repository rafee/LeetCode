/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 *
 * https://leetcode.com/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (57.53%)
 * Likes:    865
 * Dislikes: 11
 * Total Accepted:    63K
 * Total Submissions: 107.6K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * Given two strings text1 and text2, return the length of their longest common
 * subsequence.
 *
 * A subsequence of a string is a new string generated from the original string
 * with some characters(can be none) deleted without changing the relative
 * order of the remaining characters. (eg, "ace" is a subsequence of "abcde"
 * while "aec" is not). A common subsequence of two strings is a subsequence
 * that is common to both strings.
 *
 *
 *
 * If there is no common subsequence, return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: text1 = "abcde", text2 = "ace"
 * Output: 3
 * Explanation: The longest common subsequence is "ace" and its length is 3.
 *
 *
 * Example 2:
 *
 *
 * Input: text1 = "abc", text2 = "abc"
 * Output: 3
 * Explanation: The longest common subsequence is "abc" and its length is 3.
 *
 *
 * Example 3:
 *
 *
 * Input: text1 = "abc", text2 = "def"
 * Output: 0
 * Explanation: There is no such common subsequence, so the result is 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= text1.length <= 1000
 * 1 <= text2.length <= 1000
 * The input strings consist of lowercase English characters only.
 *
 *
 */

package leetcode

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	solve := make([][]int, l1+1)

	solve[l1] = make([]int, l2+1)
	for i := l1 - 1; i >= 0; i-- {
		solve[i] = make([]int, l2+1)
		for j := l2 - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				solve[i][j] = (solve[i+1][j+1] + 1)
			} else {
				solve[i][j] = max(solve[i+1][j], solve[i][j+1])
			}
		}
	}
	return solve[0][0]
}

// @lc code=end
