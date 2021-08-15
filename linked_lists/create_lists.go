package linkedlists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(arr []int) *ListNode {
	myList := &ListNode{}
	prev := &ListNode{}

	for i, v := range arr {
		curr := &ListNode{}
		if i == 0 {
			myList.Val = v
			prev = myList
		} else {
			curr.Val = v
			prev.Next = curr
			prev = curr
		}
	}

	return myList
}

func PrintList(head *ListNode) {
	a := head
	s := ""
	for a != nil {
		s += fmt.Sprintf(" %+v", a.Val)
		a = a.Next
	}

	fmt.Println(s)
}

func DeleteList(head *ListNode, key int) *ListNode {
	prev, curr := head, head.Next
	for prev.Val == key {
		prev = curr
		curr = curr.Next
	}

	h := prev
	for curr != nil {
		for curr.Val == key {
			curr = curr.Next
		}
		prev.Next = curr
		prev = curr
		curr = curr.Next
	}

	return h
}
