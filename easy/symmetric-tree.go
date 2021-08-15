package easy

func IsSymmetric(root *TreeNode) bool {
	if root != nil {
		return isSameTree(root.Left, root.Right)
	}

	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil {
		if q != nil {
			if p.Val == q.Val {
				left := isSameTree(q.Right, p.Left)
				if left {
					return isSameTree(q.Left, p.Right)
				}
			}
		}

		return false
	}

	return q == nil
}
