// main.go
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.Get("/getRequest", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("请求路径:" + path)
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
