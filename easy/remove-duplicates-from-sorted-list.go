package easy

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	a := head
	for a.Next != nil {
		if a.Val == a.Next.Val {
			a.Next = a.Next.Next
			continue
		}
		if a.Next != nil {
			a = a.Next
		}
	}
	return head
}
