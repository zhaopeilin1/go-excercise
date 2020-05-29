package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine，i:", i)
		}
		fmt.Println("子goroutine结束")
		ch1 <- true
	}()
	b := <-ch1
	fmt.Println(b)
	fmt.Println("main over")
}
