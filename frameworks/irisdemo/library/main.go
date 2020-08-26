// main.go
package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
	//"github.com/kataras/iris/context"
	"github.com/blevesearch/bleve/mapping"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	mapping := bleve.NewIndexMapping()
	index, _ := bleve.New("example.bleve", mapping)
	// app.Get("/get", func(context context.Context) {
	// 	path := context.Path()
	// 	app.Logger().Info(path)
	// 	result := search()
	// 	result.String()
	// 	fmt.Println(result)+ result.String()
	// 	context.WriteString("请求路径:" + path)
	// })

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		fmt.Println("ok")
		result := search(mapping, index)
		ctx.JSON(iris.Map{"message": result.String()})
	})
	// result := search()

	app.Listen(":8000")
	//app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

func search(mapping *mapping.IndexMappingImpl, index bleve.Index) *bleve.SearchResult {
	// open a new index
	// mapping := bleve.NewIndexMapping()
	// index, _ := bleve.New("example.bleve", mapping)

	// index some data
	// err = index.Index(identifier, "title is good")
	index.Index("id", "aa")
	index.Index("ok", "title is good")

	// search for some text
	query := bleve.NewMatchQuery("good")
	search := bleve.NewSearchRequest(query)
	searchResults, _ := index.Search(search)
	return searchResults
}
