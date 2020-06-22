package main

import (
	"fmt"
	"os"
)

func main() {
	var str, seq string

	for i := 1; i < len(os.Args); i++ {
		str += seq + os.Args[i]
		seq = " "
		//for 初始化语句;判断条件;循环动作
		//for i:=0;i<len(xx);i++
		//每个部分都可以省略，for {}无限循环

	}
	fmt.Println(str)
}
