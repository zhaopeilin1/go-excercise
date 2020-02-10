package tree

import (
	"fmt"
	//"strings"
	"reflect"
)

type TrieNode struct {
	//以该字母为结尾的单词数
	Count int
	Next  []*TrieNode
}

type CTrie struct {
	//压缩字典树
	Val  string
	Next []*CTrie
}

const (
	Aint int = int(rune('a'))
)

//插入，删除，查找，遍历输出

func (root *TrieNode) Insert(str string) {
	//插入一个单词,有几种情况。1、整个单词没有和别的共同的前缀，作为一个节点。
	//2、单词和别的单词有共同前缀，又分两种，直接到结尾，另一种有分叉

	p := root
	for _, by := range str {
		val := int(by) - Aint
		if p.Next == nil {
			p.Next = make([]*TrieNode, 26)
			p.Next[val] = &TrieNode{1, make([]*TrieNode, 26)}
		} else if p.Next[val] == nil {
			p.Next[val] = &TrieNode{1, make([]*TrieNode, 26)}
		} else {
			p.Next[val].Count = p.Next[val].Count + 1
		}
		p = p.Next[val]
	}
}

func (root *TrieNode) Delete(str string) {
	//删除一个单词
	p := root
	for _, by := range str {
		val := int(by) - Aint
		if p.Next == nil {
			p.Next = make([]*TrieNode, 26)
			p.Next[val] = &TrieNode{1, make([]*TrieNode, 26)}
		} else if p.Next[val] == nil {
			p.Next[val] = &TrieNode{1, make([]*TrieNode, 26)}
		} else {
			p.Next[val].Count = p.Next[val].Count + 1
		}
		p = p.Next[val]
	}
}

func (root *TrieNode) Search(str string) bool {
	p := root.Next
	for _, ru := range str {
		if p != nil && p[int(ru)-Aint] != nil {
			p = p[int(ru)-Aint].Next
		} else {
			return false
		}
	}
	//查找一个单词
	return true
}

func (root *TrieNode) Print() {
	//打印输出
	for index, children := range root.Next {
		if children != nil {
			val := Aint + index
			fmt.Println(string(rune(val)), children.Count)
			children.Print()
		}
	}
}

func (root *CTrie) Insert(str string) {
	//插入一个单词,有几种情况。1、整个单词没有和别的共同的前缀，作为一个节点。
	//2、单词和别的单词有共同前缀，又分两种，直接到结尾，另一种有分叉
	p := root
	index := 0
	for index < len([]rune(str)) {
		by := []rune(str)[index]
		//按首字母开头排序从a-z的26字母的数组。
		charIndex := getIndexFromRune(by)
		if p.Next == nil {
			//直接保存后续的词语后缀
			p.Next = make([]*CTrie, 27)
			p.Next[charIndex] = &CTrie{string([]rune(str)[index:]), make([]*CTrie, 27)}
			break
		} else if p.Next[charIndex] == nil {
			//直接保存后续的词语后缀
			p.Next[charIndex] = &CTrie{string([]rune(str)[index:]), make([]*CTrie, 27)}
			break
		} else {
			//2、单词和别的单词有共同前缀，又分三种，直接到结尾，一种有中途分叉，一种新来的是原有的子串
			//p.Next[val].Count = p.Next[val].Count + 1
			//有共同前缀，看前缀匹配的长度
			node := p.Next[charIndex]
			recordStr := node.Val
			//计算两个字符串的共同前缀，返回共同长度n,是否等于s1长度，是否等于s2长度。
			n, b1, b2 := calSamePrefix(recordStr, string([]rune(str)[index:]))
			//从n处截断，进行分叉。n1为record的分叉，n2为str的分叉。
			//apple,able -> a,pple,ble
			//处理本节点的字符串值。如 abc 和 abd有公共前缀 ab,就把abc的变为ab
			node.Val = string([]rune(recordStr)[0:n])
			if node.Next == nil {
				node.Next = make([]*CTrie, 27)
				if b1 && b2 {
					//node.Next = &CTrie{}
				} else if b1 {
					//不用做什么，下一次迭代处理剩下的字符串
					//node.Next[] = &CTrie{string([]rune(str)[n:]), make([]*CTrie, 26)}
				} else if b2 {
					//把原来的进行分裂成两个节点后，手动创建子节点，因为没有下一次迭代了。
					node.Next[charIndex] = &CTrie{string([]rune(recordStr)[n:]), make([]*CTrie, 27)}
				} else {
					//把原来的进行分裂成两个节点后，手动创建子节点
					node.Next[charIndex] = &CTrie{string([]rune(recordStr)[n:]), make([]*CTrie, 27)}
				}
			} else {
				//如果next不为空
				//插入中间的节点,中间节点 a->b a->c->b
				//abc->d   :  ab->c->d
				if b1 && b2 {
					//node.Next = &CTrie{}
				} else if b1 {
					//不用做什么，下一次迭代处理剩下的字符串
					//node.Next[] = &CTrie{string([]rune(str)[n:]), make([]*CTrie, 26)}
				} else if b2 {
					//完全匹配模式字符串。abc 和ab的情况原来的要中间插入一个节点
					grandChildren := node.Next
					children := &CTrie{string([]rune(recordStr)[n:]), grandChildren}
					node.Next = make([]*CTrie, 27)
					node.Next[getIndex(string([]rune(recordStr)[n:]))] = children
				} else {
					//把原来的进行分裂成两个节点后，手动创建子节点
					grandChildren := node.Next
					children := &CTrie{string([]rune(recordStr)[n:]), grandChildren}
					node.Next = make([]*CTrie, 27)
					node.Next[getIndex(string([]rune(recordStr)[n:]))] = children
				}
			}
			index += n
		}
		p = p.Next[charIndex]
	}
}

func (root *CTrie) InsertSuffix(str string) {
	//所有后缀，插入压缩字典树。
	str = str + "$"
	list := []rune(str)
	for index := len(list) - 1; index >= 0; index-- {
		root.Insert(string([]rune(str)[index:]))
	}
}

func (root *CTrie) Print() {
	//打印输出
	for _, children := range root.Next {
		if children != nil {
			//val := Aint + index
			fmt.Println(children.Val)
			children.Print()
		}
	}
}

func getIndex(str string) int {
	list := []rune(str)
	return getIndexFromRune(list[0])
}
func getIndexFromRune(ru rune) int {
	charIndex := int(ru) - Aint
	if charIndex == -61 {
		return 26
	} else {
		return charIndex
	}
}

func calSamePrefix(s1, s2 string) (int, bool, bool) {
	//计算两个字符串的共同前缀，返回共同长度n,是否等于s1长度，是否等于s2长度。
	list1 := []rune(s1)
	list2 := []rune(s2)
	l1 := len(list1)
	l2 := len(list2)
	var min int
	if l1 <= l2 {
		min = l1
	} else {
		min = l2
	}
	same := 0
	for i := 0; i < min; i++ {
		if list1[i] == list2[i] {
			same += 1
		} else {
			return same, same == l1, same == l2
		}
	}
	return same, same == l1, same == l2
}

func (root *CTrie) Search(str string) bool {
	//后缀树匹配模式
	p := root.Next
	list := []rune(str)
	l := len(list)
	index := 0

	for index < l {
		ru := list[index]
		if p != nil && p[getIndexFromRune(ru)] != nil {
			node := p[getIndexFromRune(ru)]
			recordStr := node.Val
			if len(recordStr) >= l-index {
				//后缀足够长的情况。看模式是否储存的字符串的子串
				return reflect.DeepEqual([]rune(recordStr)[:l-index], list[index:])
			} else {
				//后缀不够长的情况。 看储存的子串是否模式的子串，如果是，用子节点继续匹配
				if reflect.DeepEqual([]rune(recordStr), list[index:index+len(recordStr)]) {
					index += len(recordStr)
				} else {
					return false
				}
			}
			p = p[getIndexFromRune(ru)].Next
		} else {
			return false
		}
	}
	return true

	for _, ru := range str {
		if p != nil && p[getIndexFromRune(ru)] != nil {
			p = p[getIndexFromRune(ru)].Next
		} else {
			return false
		}
	}
	//查找一个单词
	return true
}
