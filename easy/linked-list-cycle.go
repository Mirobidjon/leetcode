package easy

func HasCycle(head *ListNode) bool {
	var m = make(map[*ListNode]int)
	for head != nil {
		_, ok := m[head]
		if ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}
