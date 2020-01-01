package leetcode83

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) say() {
	fmt.Print(n.Val)
	if n.Next != nil {
		fmt.Print("->")
		n.Next.say()
	} else {
		fmt.Println()
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	next := head.Next
	if next != nil && next.Val == head.Val {
		head.Next = next.Next
		deleteDuplicates(head)
	} else if next != nil && next.Val != head.Val {
		deleteDuplicates(head.Next)
	}
	return head
}
