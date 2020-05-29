package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch2 := make(chan int, 5)
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 100
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 200
	ch2 <- 300
	ch2 <- 400
	ch2 <- 500
	fmt.Println(len(ch2), cap(ch2))
	d := <-ch2
	fmt.Println(d)
	ch2 <- 600

	fmt.Println("----")
	ch3 := make(chan string, 4)
	go sendData(ch3)
	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("read over", ok)
			break
		}
		fmt.Println("read data:", v)
	}
}
func sendData(ch chan string) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Println("子goroutine写入n个数据，", i)
	}
}
