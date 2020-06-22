package leetcode104

import "testing"

var data = []struct {
	in  *TreeNode
	out int
}{
	{&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}, 3},
	{&TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}, 3},
}

func TestAll(t *testing.T) {
	for _, d := range data {
		r := maxDepth(d.in)
		if r != d.out {
			t.Error(d.out, r)
		}

	}

}
