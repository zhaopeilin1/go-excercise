package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径", context.FullPath())
		context.Writer.Write([]byte("\nhello,gin\n"))
	})

	if err := engine.Run(":8090"); err != nil {
		log.Fatal(err.Error)
	}
}
