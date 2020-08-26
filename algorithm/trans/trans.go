package main

import "fmt"

func main() {

	list := []int{}

	ch1 := make(chan int, 10)
	//用协程并发遍历？
	for i := 0; i < 100; i++ {
		list = append(list, i)
	}

	for i := 0; i < len(list); i++ {
		ch1 <- i
		fmt.Println(list[i])
	}

}
