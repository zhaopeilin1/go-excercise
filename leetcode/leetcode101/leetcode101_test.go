package leetcode101

import "testing"

var data = []struct {
	in  *TreeNode
	out bool
}{
	{&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}, true},
	{&TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}, false},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := isSymmetric(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}

	}

}
