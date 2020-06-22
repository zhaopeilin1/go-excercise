package leetcode104

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		l := maxDepth(root.Left)
		r := maxDepth(root.Right)
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
