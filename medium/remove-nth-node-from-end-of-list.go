package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	h, length := head, 0
	for h != nil {
		length++
		h = h.Next
	}

	var prev, curr *ListNode
	prev = head
	if head.Next != nil {
		curr = head.Next
	}

	for i := 0; i < length-n; i++ {
		prev = curr
		curr = curr.Next
	}

	if curr != nil && curr.Next != nil {
		prev.Next = curr.Next
	}

	return head
}
