package main

import (
	"fmt"
)

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

func isEndNode(t *Tree) bool {
	return (t.left == nil) && (t.right == nil)
}
func isTargetNode(t *Tree) bool {
	return (t.left != nil && isEndNode(t.left)) || (t.right != nil && isEndNode(t.right))
}

func preOrder(t *Tree, fun func(*Tree)) {
	if t == nil {
		return
	}
	fun(t)
	preOrder(t.left, fun)
	preOrder(t.right, fun)
}

func main() {
	t1 := &Tree{1, nil, nil}
	t2L := &Tree{2, nil, nil}
	t2R := &Tree{2, nil, nil}
	t3L := &Tree{3, nil, nil}
	t3R := &Tree{3, nil, nil}
	t1.left = t2L
	t1.right = t2R
	t2L.left = t3L
	t2R.right = t3R
	sum := 0

	preOrder(t1, func(tt *Tree) {
		if isTargetNode(tt) {
			sum += tt.value
		}
	})
	fmt.Println(sum)

}
