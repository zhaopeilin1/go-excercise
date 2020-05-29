// main
package main

import (
	"github.com/kataras/iris"
)

func main() {
	//创建app结构体对象
	app := iris.New()
	app.Run(iris.Addr(":7999"), iris.WithoutServerError(iris.ErrServerClosed))

}
