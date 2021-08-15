package medium

func SwapPairs(head *ListNode) *ListNode {
	var (
		curr, prev, next, h *ListNode
		have                bool
	)
	prev = head
	if head != nil && head.Next != nil {
		curr = head.Next
		next = curr.Next
		have = true
		h = curr
	} else {
		h = head
	}

	for have {
		curr.Next = prev
		prev.Next = next
		if next != nil && next.Next != nil {
			have = true
			prev.Next = next.Next
			prev = next
			curr = next.Next
			next = curr.Next
		} else {
			have = false
		}
	}

	return h
}
