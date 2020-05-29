package main

import "fmt"

func main() {
	ch1 := make(chan int) //单向 写
	//	ch2 := make(chan<- int) //单向 写
	//	ch3 := make(<-chan int) //单向 写
	//ch2 <- 100

	//data := <-ch2
	// go fun1(ch1)
	go fun1(ch1)
	data := <-ch1
	fmt.Println("read from fun1", data)
	go fun2(ch1)
	ch1 <- 22

}
func fun1(ch chan<- int) {
	ch <- 100
	fmt.Println("fun1 over")
}

func fun2(ch <-chan int) {
	d := <-ch
	fmt.Println("fun2 read", d)
}
