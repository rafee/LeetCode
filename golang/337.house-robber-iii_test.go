/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 *
 * https://leetcode.com/problems/house-robber-iii/description/
 *
 * algorithms
 * Medium (49.36%)
 * Likes:    2283
 * Dislikes: 49
 * Total Accepted:    142K
 * Total Submissions: 284.1K
 * Testcase Example:  '[3,2,3,null,3,null,1]'
 *
 * The thief has found himself a new place for his thievery again. There is
 * only one entrance to this area, called the "root." Besides the root, each
 * house has one and only one parent house. After a tour, the smart thief
 * realized that "all houses in this place forms a binary tree". It will
 * automatically contact the police if two directly-linked houses were broken
 * into on the same night.
 *
 * Determine the maximum amount of money the thief can rob tonight without
 * alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3,null,3,null,1]
 *
 * ⁠    3
 * ⁠   / \
 * ⁠  2   3
 * ⁠   \   \
 * ⁠    3   1
 *
 * Output: 7
 * Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
 *
 * Example 2:
 *
 *
 * Input: [3,4,5,1,3,null,1]
 *
 * 3
 * ⁠   / \
 * ⁠  4   5
 * ⁠ / \   \
 * ⁠1   3   1
 *
 * Output: 9
 * Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
 *
 */

package golang

import (
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robBinary(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := robBinary(tt.args.root)
			if got != tt.want {
				t.Errorf("robBinary() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("robBinary() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
