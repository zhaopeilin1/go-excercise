package chans

import "testing"

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestAll(t *testing.T) {
	strs := []string{"1", "2", "3", "4", "waat://furhurterwe.geds"}
	m := testUrls(mockWebsiteChecker, strs)
	t.Log(m)

}
