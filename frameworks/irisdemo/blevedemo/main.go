// main.go
package main

import (
	"fmt"

	"github.com/RivenZoo/blevejieba"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

func main() {

	mapping := bleve.NewIndexMapping()

	index, _ := bleve.Open("example.bleve")

	result := search(mapping, index)
	fmt.Println(result)

}

func search(mapping *mapping.IndexMappingImpl, index bleve.Index) *bleve.SearchResult {

	blevejieba.NewOptions()

	//index.Index("id", "因为是使用dnf来升级，只要软件源又新又快，升级是一件很简单的事。但是，阿里云官方镜像不够清华镜像更新快。2020-04-28，清华同步了32的文件夹，阿里云显示失败。我马上切换使用清华源。")
	//index.Index("ok", "title is good")

	// search for some text
	query := bleve.NewMatchQuery("升级")
	search := bleve.NewSearchRequest(query)
	searchResults, _ := index.Search(search)
	return searchResults
}
