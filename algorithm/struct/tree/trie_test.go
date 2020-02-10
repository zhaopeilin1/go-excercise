package tree

import "testing"

func TestTrieInsert(t *testing.T) {
	root := &TrieNode{0, nil}
	root.Insert("apple")
	//root.Print()
	root.Insert("alis")
	root.Insert("application")
	root.Print()
}

func TestTrieSearch(t *testing.T) {
	root := &TrieNode{0, nil}
	root.Insert("apple")
	//root.Print()
	root.Insert("alis")
	root.Insert("application")

	b1 := root.Search("apple")
	b2 := root.Search("aliss")
	b3 := root.Search("baby")
	t.Log(b1, b2, b3)

}

func TestCTrieInsert(t *testing.T) {
	root := &CTrie{"", nil}
	root.Insert("abc")
	root.Print()
	root.Insert("abcdd")
	root.Insert("application")
	root.Print()
}

func TestCTrieInsertSuffix(t *testing.T) {
	root := &CTrie{"", nil}
	root.InsertSuffix("xmadamyx")
	root.Print()
}

func TestSuffixTrieSearch(t *testing.T) {
	root := &CTrie{"", nil}
	root.InsertSuffix("xmadamyx")
	//root.Print()
	b1 := root.Search("xmadamyx")
	b2 := root.Search("xmad")
	b3 := root.Search("amyx")
	b4 := root.Search("madamy")
	b5 := root.Search("madamya")
	t.Log(b1, b2, b3, b4, b5)
}
