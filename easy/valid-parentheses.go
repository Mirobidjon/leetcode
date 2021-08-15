package easy

func IsValid(s string) bool {
	var a []rune
	for _, v := range s {
		switch v {
		case '}':
			if len(a) > 0 && a[len(a)-1] == '{' {
				a = a[:len(a)-1]
			} else {
				return false
			}
		case ')':
			if len(a) > 0 && a[len(a)-1] == '(' {
				a = a[:len(a)-1]
			} else {
				return false
			}
		case ']':
			if len(a) > 0 && a[len(a)-1] == '[' {
				a = a[:len(a)-1]
			} else {
				return false
			}
		default:
			a = append(a, v)
		}
	}

	return len(a) == 0
}
