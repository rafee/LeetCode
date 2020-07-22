/*
 * @lc app=leetcode id=112 lang=typescript
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function hasPathSum({ root, sum }: { root: TreeNode | null; sum: number }): boolean {
    if (root === null) {
        return false
    }
    if (root.left === null && root.right === null) {
        return root.val === sum
    }
    return (
        hasPathSum({ root: root.left, sum: sum - root.val }) ||
        hasPathSum({ root: root.right, sum: sum - root.val })
    )
}
// @lc code=end
