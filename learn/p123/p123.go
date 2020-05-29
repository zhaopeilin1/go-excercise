package main

import (
	"fmt"
)

func main() {
	/*
		通道关闭
	*/
	ch1 := make(chan int)
	go send(ch1)

	for {
		data, ok := <-ch1
		if ok {
			fmt.Println("读取到数据", data)
		} else {
			fmt.Println("读完了，没货了")
			break
		}
	}

}
func send(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
