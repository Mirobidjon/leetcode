package easy

func UniqueOccurrences(arr []int) bool {
	var (
		values []int
		used   = make([]bool, len(arr))
	)

	for i, v := range arr {
		if !used[i] {
			count := 1
			// go func() {
			for j := i + 1; j < len(arr); j++ {
				if v == arr[j] {
					used[j] = true
					count++
				}
			}
			values = append(values, count)
			// }()
		}
	}

	for i, v := range values {
		for j := i + 1; j < len(values); j++ {
			if v == values[j] {
				return false
			}
		}
	}

	return true
}
