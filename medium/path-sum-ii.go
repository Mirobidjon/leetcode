package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	var (
		arr    [][]int
		toLeaf []int
	)

	if root != nil {
		return dfs(root, 0, targetSum, arr, toLeaf)
	}

	return arr
}

func dfs(r *TreeNode, sum, targetSum int, arr [][]int, toLeaf []int) [][]int {
	if r.Right == nil && r.Left == nil {
		if targetSum == sum+r.Val {
			cp := append([]int{}, toLeaf...)
			cp = append(cp, r.Val)
			return append(arr, cp)
		}

		return arr
	}

	if r.Right != nil && r.Left != nil {
		arr = dfs(r.Right, sum+r.Val, targetSum, arr, append(toLeaf, r.Val))
		arr = dfs(r.Left, sum+r.Val, targetSum, arr, append(toLeaf, r.Val))
		return arr
	}

	if r.Right == nil {
		return dfs(r.Left, sum+r.Val, targetSum, arr, append(toLeaf, r.Val))
	}

	return dfs(r.Right, sum+r.Val, targetSum, arr, append(toLeaf, r.Val))
}
