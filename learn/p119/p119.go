package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		waitgroup
		add()设置等待组当中要执行的子 routine
		要使用 Wait，让主的等待
		Done()让等待组当中的counter计数哭的值减1，add(-1)
	*/
	wg.Add(3)
	go fun1()
	go fun2()

	fmt.Println("main进入阻塞状态，等待wg中的子routine结束")
	wg.Wait() //main goroutine等待阻塞

	fmt.Println("main解除阻塞，")

}
func fun1() {
	for i := 1; i < 3; i++ {
		fmt.Println("fun1()", i)
	}
	//让wg等待组当中的counter数值减一
	wg.Done()
}

func fun2() {
	defer wg.Done()
	for j := 1; j < 3; j++ {
		fmt.Println("fun2()", j)
	}
	//让wg等待组当中的counter数值减一

}
