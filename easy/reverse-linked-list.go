package easy

func ReverseList(head *ListNode) *ListNode {
	prev, curr := head, head.Next
	prev.Next = nil
	var a *ListNode
	for curr != nil {
		a = curr.Next
		curr.Next = prev
		prev = curr
		curr = a
	}

	return prev
}
