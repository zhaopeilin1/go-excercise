package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/**
这次使用ioutil工具，一次读取整个文件
再按行切分，同时，工具只接收string作为读取参数
调整文件格式
*/
func main() {
	counts := make(map[string]int)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(" not file")
	} else {
		for _, arg := range args {
			countsFile(arg, counts)

		}

	}
	outputMap(counts)

}
func outputMap(counts map[string]int) {
	for content, counts := range counts {
		if counts > 1 {
			fmt.Println(counts, content)
		}
	}

}
func countsFile(file string, counts map[string]int) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(" open file error", file, err)
		return
	}
	for _, str := range strings.Split(string(bytes), "\n") {
		counts[str]++
	}

}
