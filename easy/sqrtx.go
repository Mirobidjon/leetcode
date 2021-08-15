package easy

func MySqrt(x int) int {
	curr := 1.0
	prev := 0.0

	for count := 1; count <= 20; count++ {
		prev = curr
		curr = prev - (prev*prev-float64(x))/(2*prev)
		if curr == prev {
			break
		}
	}
	return int(curr)
}
