/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 *
 * https://leetcode.com/problems/reverse-string/description/
 *
 * algorithms
 * Easy (64.93%)
 * Likes:    1164
 * Dislikes: 644
 * Total Accepted:    629.1K
 * Total Submissions: 952.8K
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * Write a function that reverses a string. The input string is given as an
 * array of characters char[].
 *
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 *
 * You may assume all the characters consist of printable ascii
 * characters.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["h","e","l","l","o"]
 * Output: ["o","l","l","e","h"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["H","a","n","n","a","h"]
 * Output: ["h","a","n","n","a","H"]
 *
 *
 *
 *
 */

package golang

// recursively solve reverseString

// @lc code=start
func reverseString(s []byte) {
	lenS := len(s)

	for i := 0; i < lenS/2; i++ {
		s[i], s[lenS-i-1] = s[lenS-i-1], s[i]
	}
}

// @lc code=end