/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return validate(root, root)
}

func validate(root1, root2 *TreeNode) bool {
    if root1 == nil || root2 == nil {
        return root1 == root2
    }
    
    if root1.Val != root2.Val {
        return false
    }
    return validate(root1.Left,root2.Right) && validate(root1.Right,root2.Left)
}
