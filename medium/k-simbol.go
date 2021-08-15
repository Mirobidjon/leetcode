package medium

import "fmt"

func KthGrammar(n int, k int) int {
	var s string = "0"
	for i := 1; i < n; i++ {
		a := len(s) / 2
		b := len(s)
		for j := a; j < b; j++ {
			if s[j] == '0' {
				s += "01"
			} else {
				s += "10"
			}
		}

		if i == 1 {
			s = s[1:]
		}
	}

	fmt.Println(s)
	if s[k-1] == '0' {
		return 0
	}

	return 1
}

/*
0
01
0110
0110 1001
0110 1001 1001 0110
0110 1001 1001 0110 1001 0110 0110 1001
0110 1001 1001 0110 1001 0110 0110 1001 1001 0110 0110 1001 0110 1001 1001 0110
0110 1001 1001 0110 1001 0110 0110 1001 1001 0110 0110 1001 0110 1001 1001 0110 1001 011001101001011010011001011001101001100101101001011001101001

*/
