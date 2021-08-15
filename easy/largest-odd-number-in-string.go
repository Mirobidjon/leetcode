package easy

func LargestOddNumber(num string) string {
	a := int(num[len(num)-1] - '0')
	for a%2 == 0 && len(num) > 0 {
		num = num[:len(num)-1]
		// fmt.Println(num)
		if len(num) > 0 {
			a = int(num[len(num)-1] - '0')
		}
	}

	return num
}
