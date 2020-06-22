package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go sendData(ch1)
	for v := range ch1 {
		fmt.Println("read data", v)

	}
	fmt.Println("main over")
}
func sendData(ch1 chan int) {
	defer close(ch1)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch1 <- i
	}
	//close(ch1)
}
