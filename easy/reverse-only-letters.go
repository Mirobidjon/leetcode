package easy

func ReverseOnlyLetters(s string) string {
	var (
		a = make(map[int]rune)
	)
	for i, v := range s {
		if !(('a' <= v && v <= 'z') || ('A' <= v && v <= 'Z')) {
			a[i] = v
			s = s[:i] + s[i+1:]
		}
	}

	s = Reverse(s)
	for i, v := range a {
		s = s[:i] + string(v) + s[i:]
	}

	return s
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
