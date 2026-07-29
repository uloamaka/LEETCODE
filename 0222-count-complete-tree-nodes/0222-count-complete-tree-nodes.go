/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeftHeight(root *TreeNode) int {
    count := 0
    for root != nil {
        root = root.Left
        count++
    }
    return count
}

func findRightHeight(root *TreeNode) int {
    count := 0
    for root != nil {
        root = root.Right
        count++
    }
    return count
}

func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    l := findLeftHeight(root)
    r := findRightHeight(root)
    
    // If the left and right depths are equal, it's a perfect binary tree.
    if l == r {
        // (1 << l) is equivalent to 2^l. 
        return (1 << l) - 1 
    }
    
    // Otherwise, recursively count the left and right subtrees.
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}