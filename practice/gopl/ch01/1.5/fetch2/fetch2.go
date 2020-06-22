package main

import (
	"fmt"
	"io"

	// "io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	//使用net/http Get函数获取网页内容，用ioutil.ReadAll读取网页返回内容
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		//使用io.Copy方法
		fmt.Printf("http status %v \n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
	}
}
