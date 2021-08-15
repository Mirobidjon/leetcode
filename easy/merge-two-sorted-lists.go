package easy

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		s *ListNode
		h *ListNode
	)

	for l1 != nil || l2 != nil {
		a := &ListNode{}
		if (l1 != nil && l2 != nil && l1.Val > l2.Val) || (l2 != nil && l1 == nil) {
			a.Val = l2.Val
			l2 = l2.Next
		} else {
			a.Val = l1.Val
			l1 = l1.Next
		}

		if s != nil {
			s.Next = a
			s = a
		} else {
			h = a
			s = h
		}
	}

	return h
}
