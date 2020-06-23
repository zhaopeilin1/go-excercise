package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	c := make(chan int, 1000)
	fmt.Println(len(c))
	for i := 0; i < 1000; i++ {
		c <- i
		wg.Add(1)
	}

	fmt.Println(len(c))

	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c, wg)
	}
	wg.Wait()
	//time.Sleep(time.Millisecond * 3000)
	/*
		for {
			select {
			//当c通道满了时，此时操作阻塞。
			case c <- rand.Int():
				// 可选的代码在这里
				fmt.Println("send rand int")

			case <-time.After(time.Millisecond * 100):
				//fmt.Println("timed out")

				//default:
				// 这里可以留空以静默删除数据
				//fmt.Println("dropped")
			}

			time.Sleep(time.Millisecond * 50)
		}
	*/
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int, wg sync.WaitGroup) {
	for {
		//fmt.Println(len(c))
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 10)
		wg.Done()
	}
}
