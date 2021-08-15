package easy

func SumOfLeftLeaves(root *TreeNode) int {
	if root != nil {
		return sumOfLeftLeavesHelper(root, 0, false)
	}

	return 0
}

func sumOfLeftLeavesHelper(r *TreeNode, s int, h bool) int {
	if r.Left == nil && r.Right == nil {
		if h {
			return r.Val
		}

		return 0
	}

	if r.Left != nil && r.Right != nil {
		return sumOfLeftLeavesHelper(r.Left, s, true) + sumOfLeftLeavesHelper(r.Right, s, false)
	}

	if r.Left != nil {
		return sumOfLeftLeavesHelper(r.Left, s, true)
	}

	return sumOfLeftLeavesHelper(r.Right, s, false)
}
