/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // 1 (current node) + count of left branch + count of right branch
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}