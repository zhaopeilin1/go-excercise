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
