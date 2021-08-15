package medium

func RotateRight(head *ListNode, k int) *ListNode {
	n := 0
	h := head
	for h != nil {
		n++
		h = h.Next
	}

	if n == 1 || n == k || k == 0 || n == 0 || k%n == 0 {
		return head
	}

	k = k % n
	h = head
	first := h

	for i := 1; i < n-k; i++ {
		h = h.Next
	}

	first = h.Next
	h.Next = nil

	h = first

	for h.Next != nil {
		h = h.Next
	}

	h.Next = head
	return first
}
