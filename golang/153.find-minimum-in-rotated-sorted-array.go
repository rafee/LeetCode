/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.94%)
 * Likes:    1770
 * Dislikes: 217
 * Total Accepted:    399.7K
 * Total Submissions: 900.4K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 *
 * Find the minimum element.
 *
 * You may assume no duplicate exists in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,4,5,1,2]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,5,6,7,0,1,2]
 * Output: 0
 *
 *
 */

package golang

// @lc code=start
func findMinI(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < nums[end] {
			end = mid
		} else if nums[mid] > nums[end] {
			start = mid + 1
		}
	}
	return nums[start]
}

// @lc code=end
