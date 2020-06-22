package main

import (
	"fmt"
)

func main() {
	/*
		channel 类型
	*/
	var a chan int
	fmt.Printf("%T,%v\n", a, a)
	if a == nil {
		fmt.Println("nil 要创建")
		a = make(chan int)
		fmt.Println(a)
	}
	test1(a)

	a <- 1
	b := <-a
	fmt.Println(b)

}
func test1(ch chan int) {
	fmt.Printf("%T,%v\n", ch, ch)
}
