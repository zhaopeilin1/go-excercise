package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine开始")
		data := <-ch1
		time.Sleep(3 * time.Second)
		fmt.Println("从ch中读了", data)

	}()
	ch1 <- 11

	fmt.Println("main over")
}
