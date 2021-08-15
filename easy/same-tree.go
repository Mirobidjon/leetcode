package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil {
		if q != nil {
			if p.Val == q.Val {
				left := IsSameTree(q.Left, p.Left)
				if left {
					return IsSameTree(q.Right, p.Right)
				}
			}
		}

		return false
	}

	return q == nil
}
