package medium

func SortList(head *ListNode) *ListNode {
	var (
		// s *ListNode
		h *ListNode
		n int
	)

	h = head
	for h != nil {
		n++
		h = h.Next
	}

	h = head

	// for i:=0; i<n; i++{
	// 	temp := &ListNode{}
	// 	min =

	// }
	return h
}
