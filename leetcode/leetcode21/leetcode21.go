package leetcode21

import (
	"fmt"
)

func (n *ListNode) say() {
	fmt.Print(n.Val)
	if n.Next != nil {
		fmt.Print("->")
		n.Next.say()
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3 := &ListNode{l1.Val, nil}
			mergeTwoLists2(l1.Next, l2, l3)
			return l3
		} else {
			l3 := &ListNode{l2.Val, nil}
			mergeTwoLists2(l1, l2.Next, l3)
			return l3
		}
	} else if l1 != nil {
		return l1
	} else if l2 != nil {
		return l2
	} else {
		return nil
	}
}
func mergeTwoLists2(l1 *ListNode, l2 *ListNode, l3 *ListNode) {
	if l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3.Next = &ListNode{l1.Val, nil}
			mergeTwoLists2(l1.Next, l2, l3.Next)
		} else {
			l3.Next = &ListNode{l2.Val, nil}
			mergeTwoLists2(l1, l2.Next, l3.Next)
		}
	} else if l1 != nil {
		l3.Next = l1
	} else if l2 != nil {
		l3.Next = l2
	}
}
