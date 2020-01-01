package leetcode83

import "testing"

var data = []struct {
	in  *ListNode
	out *ListNode
}{
	{&ListNode{1, &ListNode{1, nil}}, &ListNode{1, nil}},
	{&ListNode{2, &ListNode{3, nil}}, &ListNode{2, &ListNode{3, nil}}},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := deleteDuplicates(d.in)
		r.say()
		d.out.say()
	}

}
