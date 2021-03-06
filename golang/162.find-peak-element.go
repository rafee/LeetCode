/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 *
 * https://leetcode.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (42.28%)
 * Likes:    1441
 * Dislikes: 1862
 * Total Accepted:    328.3K
 * Total Submissions: 767.3K
 * Testcase Example:  '[1,2,3,1]'
 *
 * A peak element is an element that is greater than its neighbors.
 *
 * Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element
 * and return its index.
 *
 * The array may contain multiple peaks, in that case return the index to any
 * one of the peaks is fine.
 *
 * You may imagine that nums[-1] = nums[n] = -∞.
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index
 * number 2.
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 1 or 5
 * Explanation: Your function can return either index number 1 where the peak
 * element is 2,
 * or index number 5 where the peak element is 6.
 *
 *
 * Note:
 *
 * Your solution should be in logarithmic complexity.
 *
 */

package golang

// @lc code=start
func findPeakElement(nums []int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (start + end) / 2
		var left, right bool
		if mid == 0 {
			left = true
		} else {
			left = nums[mid] > nums[mid-1]
		}
		if mid == len(nums)-1 {
			right = true
		} else {
			right = nums[mid] > nums[mid+1]
		}
		if left && right {
			return mid
		} else if left && !right {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

// @lc code=end
