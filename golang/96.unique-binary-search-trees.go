/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (48.83%)
 * Likes:    2694
 * Dislikes: 100
 * Total Accepted:    257.9K
 * Total Submissions: 515.1K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */

package golang

// @lc code=start
func numTrees(n int) int {
	sol := make([]int, n+1)
	sol[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			sol[i] += (sol[j-1] * sol[i-j])
		}
	}
	return sol[n]
}

// @lc code=end
