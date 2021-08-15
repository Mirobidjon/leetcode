package medium

func CountGoodNumbers(n int64) int {
	m := int64(1e9 + 7)
	if n < 3 {
		if n == 2 {
			return 16
		}

		return 5
	}

	a := binPow((n+1)/2, 5, m)
	b := binPow(n/2, 4, m)
	return int((a * b) % m)
}

func binPow(n int64, a int64, m int64) int64 {
	r := int64(1)
	for n > 0 {
		if n%2 == 1 {
			r = (r * a) % m
		}

		a = (a * a) % m
		n /= 2
	}

	return r % m
}
