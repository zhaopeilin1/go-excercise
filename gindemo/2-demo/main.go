package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		name := context.DefaultQuery("name", "小明")
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("\nhello," + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		//post表单传参
		name := context.PostForm("name")
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("\nhello," + name))
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error)
	}
}
