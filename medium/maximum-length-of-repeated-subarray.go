package medium

import "fmt"

func FindLength(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	var (
		max int
		dp  [1003][1002]int
	)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Printf("\n")
	}

	return max
}
