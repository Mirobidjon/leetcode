package medium

func DeleteDuplicatesII(head *ListNode) *ListNode {
	var (
		prev, curr *ListNode
	)

	prev = head
	if head != nil && head.Next != nil {
		curr = head.Next
		if curr == nil {
			return prev
		}

		if prev.Val == curr.Val {
			for curr != nil && prev.Val == curr.Val {
				curr = curr.Next
			}
			prev = curr
			head = curr
			if curr == nil {
				return prev
			}
			curr = curr.Next
		}
	}

	for prev.Next != nil {
		var have bool
		for curr != nil && curr.Next != nil && curr.Val == curr.Next.Val {
			have = true
			curr.Next = curr.Next.Next
		}

		if have {
			prev.Next = curr.Next
		} else {
			prev.Next = curr
		}

		curr = curr.Next
	}

	return head
}
