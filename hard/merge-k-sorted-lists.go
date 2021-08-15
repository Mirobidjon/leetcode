package hard

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	var (
		s *ListNode
		h *ListNode
		k int
	)

	n := len(lists)

	for {
		k = 0
		a := &ListNode{}
		max := 0
		for i := 0; i < n; i++ {
			if lists[i] == nil {
				k++
			} else {
				if lists[max] == nil {
					max = i
				} else {
					if lists[max].Val > lists[i].Val {
						max = i
					}
				}
			}
		}

		if k == n {
			break
		}

		a.Val = lists[max].Val
		lists[max] = lists[max].Next

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
