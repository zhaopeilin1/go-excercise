package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
	}()

	select {
	case <-ch1:
		fmt.Println("ch1 case1")
	case <-ch2:
		fmt.Println("ch2 case2")
	case <-time.After(3 * time.Second):
		fmt.Println("3 s timeout")

	}

}
