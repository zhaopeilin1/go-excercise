package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("from ch1", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("from ch2", num2)
		} else {
			fmt.Println("ch2 close")
		}
	default:
		fmt.Println("default")
	}

	fmt.Println("main over")

}

func query() {
	time.Sleep(2 * time.Second)
}

func delay1S() {
	time.Sleep(1 * time.Second)
}
