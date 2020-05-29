package main

import (
	"fmt"
	"time"
)

func main() {
	a := 10
	go func() {
		a = 20
		fmt.Println("go routine,", a)

	}()
	a = 30
	time.Sleep(1)
	fmt.Println("main a ", a)

}
