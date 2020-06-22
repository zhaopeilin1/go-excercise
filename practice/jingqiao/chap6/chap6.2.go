package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
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
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		//fmt.Println(len(c))
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}
