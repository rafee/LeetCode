/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (49.34%)
 * Likes:    496
 * Dislikes: 66
 * Total Accepted:    49.1K
 * Total Submissions: 99K
 * Testcase Example:  '13'
 *
 * Given an integer n, return 1 - n in lexicographical order.
 *
 * For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].
 *
 * Please optimize your algorithm to use less time and space. The input size
 * may be as large as 5,000,000.
 *
 */

package golang

// @lc code=start
func lexicalOrder(n int) []int {
	res := make([]int, 0)
	for i := 1; i < 10 && i <= n; i++ {
		lexicalDFS(i, n, &res)
	}
	return res
}
func lexicalDFS(st, n int, res *[]int) {
	*res = append(*res, st)
	st *= 10
	for i := 0; i < 10; i++ {
		if st+i > n {
			return
		}
		lexicalDFS(st+i, n, res)
	}
}

// @lc code=end
