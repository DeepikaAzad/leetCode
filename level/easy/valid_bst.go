func isValidBST(root *TreeNode) bool {
    return validate(root, nil, nil)
}

func validate(root *TreeNode, max, min *int) bool {
    if root == nil {
        return true
    }
    
    if (max != nil && *max <= root.Val) || (min != nil && *min >= root.Val) {
        return false
    }
    return validate(root.Left, &(root.Val), min) && validate(root.Right, max, &(root.Val))
}
