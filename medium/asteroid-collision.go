package medium

import (
	"fmt"
	"math"
)

func AsteroidCollision(A []int) []int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if j < len(A)-1 && different(A[j], A[j+1]) {
				if math.Abs(float64(A[j])) == math.Abs(float64(A[j+1])) {
					A = removeItem(A, j)
					A = removeItem(A, j)
				}
				if math.Abs(float64(A[j])) < math.Abs(float64(A[j+1])) {
					A = removeItem(A, j)
				} else {
					fmt.Printf("%+v, one %d\n\n", A, j)
					A = A[:j+1]
				}
				fmt.Println(A)
			}
		}
	}

	return A
}

func different(a, b int) bool {
	if (a > 0 && b > 0) || (a < 0 && b < 0) || (a < 0 && b > 0) {
		return false
	}
	return true
}

func removeItem(A []int, i int) []int {
	return append(A[:i], A[(i+1):]...)
}
