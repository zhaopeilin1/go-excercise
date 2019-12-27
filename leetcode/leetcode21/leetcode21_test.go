package leetcode21

import (
	"testing"
)

var data = []struct {
	in1 *ListNode
	in2 *ListNode
	out *ListNode
}{
	{&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
	},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := mergeTwoLists(d.in1, d.in2)
		r.say()
		println()
		d.out.say()
		if r != d.out {
			t.Error("expected:", d.out, " got:", r)
		}

	}

}
