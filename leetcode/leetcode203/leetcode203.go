package leetcode203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	} else if head.Val == val {
		return removeElements(head.Next, val)
	} else {
		var next *ListNode
		next = head.Next
		if next != nil {
			if next.Val == val {
				head.Next = next.Next
				removeElements(head, val)
			} else {
				removeElements(next, val)
			}
		}
		return head
	}
}
