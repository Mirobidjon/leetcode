package medium

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	var (
		from, curr, to *ListNode

		i int
	)
	h := head
	for i < left {
		i++
		from = h
		h = h.Next
	}

	curr = h
	for i < right-left {
		i++
		h = h.Next
	}

	to = h.Next
	h.Next = nil
	curr = reverseList(curr)

	from.Next = curr
	for curr != nil && curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = to
	return head
}

func reverseList(head *ListNode) *ListNode {
	var a, curr *ListNode
	prev := head
	if head != nil {
		curr = head.Next
		prev.Next = nil
		for curr != nil {
			a = curr.Next
			curr.Next = prev
			prev = curr
			curr = a
		}
	}

	return prev
}
