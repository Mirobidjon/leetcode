package easy

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root != nil {
		return dfs(root, 0, targetSum)
	}

	return false
}

func dfs(r *TreeNode, sum, targetSum int) bool {
	if r.Right == nil && r.Left == nil {
		return targetSum == sum+r.Val
	}

	if r.Right != nil && r.Left != nil {
		if !dfs(r.Right, sum+r.Val, targetSum) {
			return dfs(r.Left, sum+r.Val, targetSum)
		}

		return true
	}

	if r.Right == nil {
		return dfs(r.Left, sum+r.Val, targetSum)
	}

	return dfs(r.Right, sum+r.Val, targetSum)
}
