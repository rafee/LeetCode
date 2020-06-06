/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

package leetcode

import (
	"sort"
)

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	count := make([]int, len(nums))
	parent := make([]int, len(nums))
	max, maxIndex := 1, 0
	sort.Ints(nums)
	for i, num := range nums {
		count[i] = 1
		parent[i] = i
		for j := i - 1; j >= 0; j-- {
			if num%nums[j] == 0 && count[j] > count[i]-1 {
				count[i] = count[j] + 1
				parent[i] = j
				if count[i] > max {
					max = count[i]
					maxIndex = i
				}
			}
		}
	}
	result := make([]int, max)
	for i := max - 1; i >= 0; i-- {
		result[i] = nums[maxIndex]
		maxIndex = parent[maxIndex]
	}
	return result
}

// @lc code=end
