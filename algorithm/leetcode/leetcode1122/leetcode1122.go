package leetcode1122

import (
	"sort"
)

type IntList struct {
	arr1   []int
	orders map[int]int // attr2 value,index
}

func (p IntList) Len() int {
	return len(p.arr1)
}

func (p IntList) Less(i, j int) bool {
	//两个都在map里，一个在map里，另一个不在，两个都不在
	vi := p.arr1[i]
	vj := p.arr1[j]
	order1, exists1 := p.orders[vi]
	order2, exists2 := p.orders[vj]
	if exists1 && exists2 {
		return order1 < order2
	} else if exists1 {
		return true
	} else if exists2 {
		return false
	} else {
		return vi < vj
	}
}

func (p IntList) Swap(i, j int) {
	p.arr1[i], p.arr1[j] = p.arr1[j], p.arr1[i]
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for i, v := range arr2 {
		m[v] = i
	}
	list := IntList{arr1, m}
	sort.Sort(list)
	return list.arr1
}
