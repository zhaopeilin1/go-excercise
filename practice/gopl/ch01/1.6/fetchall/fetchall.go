package main

import (
	"fmt"
	"io"

	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//使用net/http Get函数获取网页内容，用ioutil.ReadAll读取网页返回内容
	start := time.Now()

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	/**
	  当一个goroutine尝试在一个channel上做send或者receive操作时,这个goroutine会阻塞在调用
	  处,直到另一个goroutine往这个channel里写入、或者接收值,这样两个goroutine才会继续执行
	  channel操作之后的逻辑。在这个例子中,每一个fetch函数在执行时都会往channel里发送一个值
	  (ch <­ expression),主函数负责接收这些值(<­ch)。这个程序中我们用main函数来接收所有fetch函
	  数传回的字符串,可以避免在goroutine异步执行还没有完成时main函数提前退出。
	*/
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2f seconds elapsed \n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "copy: %s: %v\n", url, err)
		os.Exit(1)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f seconds %7d  %s", secs, bytes, url)
}
