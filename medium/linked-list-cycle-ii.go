package medium

func DetectCycle(head *ListNode) *ListNode {
	var m = make(map[*ListNode]int)
	for head != nil {
		_, ok := m[head]
		if ok {
			return head
		}
		m[head] = 1
		head = head.Next
	}
	return nil
}
