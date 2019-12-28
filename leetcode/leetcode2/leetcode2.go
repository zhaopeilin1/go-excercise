package leetcode2

import (
	"fmt"
)

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{6, &ListNode{9, nil}}}}
	//l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// fmt.Println(l1)
	// fmt.Println(l2)

	l3 := addTwoNumbers(l1, l2)

	// fmt.Println(l3)
	l1.prin()
	fmt.Println()
	l2.prin()
	fmt.Println()
	l3.prin()
	fmt.Println()
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	l3 = addTwoNumbers(l1, l2)

	// fmt.Println(l3)
	l1.prin()
	fmt.Println()
	l2.prin()
	fmt.Println()
	l3.prin()

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers3(&ListNode{Val: 0}, l1, l2)
}

func addTwoNumbers3(l3 *ListNode, l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
	} else {
		var l22, l11 *ListNode
		var v11, v22 int
		if nil == l1 {
			v11 = 0
		} else {
			v11 = l1.Val
			l11 = l1.Next
		}
		if nil == l2 {
			v22 = 0
		} else {
			v22 = l2.Val
			l22 = l2.Next
		}
		r := v11 + v22 + l3.Val
		l3.Val = r % 10
		//l1,l2 相加，并把进位存下来.
		if r/10 == 0 && l11 == nil && l22 == nil {
		} else {
			l3.Next = &ListNode{Val: r / 10}
			addTwoNumbers3(l3.Next, l11, l22)
		}
	}
	return l3
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s *ListNode) prin() {
	fmt.Print(s.Val)
	if s.Next != nil {
		fmt.Print("->")
		s.Next.prin()
	}

}

/*
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
