package leetcode101

func isSymmetric(root *TreeNode) bool {
	//检查是否对称的二叉树
	if root == nil {
		return true
	}
	p := root.Left
	q := root.Right
	return isSameTree2(p, q)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else {
		if p != nil {
			if q != nil {
				return p.Val == q.Val && isSameTree2(p.Right, q.Left) && isSameTree2(p.Left, q.Right)
			} else {
				return false
			}
		} else {
			return false
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
