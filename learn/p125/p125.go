package main

import (
	"fmt"
	//"time"
)

func main() {
	/*
		单向通道
	*/
	ch1 := make(chan string)
	done := make(chan bool)
	go sendData(ch1, done)
	data := <-ch1
	fmt.Println("读到,", data)
	ch1 <- "我是主"
	fmt.Println("main over")
	//time.Sleep(1 * time.Second)
	<-done

}
func sendData(ch1 chan string, done chan bool) {
	ch1 <- "我是韩"
	d := <-ch1
	fmt.Println(d, "from main")
	done <- true
}
