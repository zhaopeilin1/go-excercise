package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int = 0
	lock    sync.Mutex
	ch      chan int = make(chan int)
)

func main() {

	fmt.Println("start")
	for i := 0; i < 5; i++ {
		w := &Worker{i}
		go w.dowork(ch)
	}
	for i := 0; i < 22; i++ {
		go send(i)
	}
	time.Sleep(time.Millisecond * 1000) // this is bad, don't do this!
	fmt.Println("done")
}

type Worker struct {
	Id int
}

func send(i int) {
	ch <- i
}

func (w Worker) dowork(c chan int) {
	for {
		r := <-c
		fmt.Printf("worker %d got %d \n", w.Id, r)
		time.Sleep(time.Millisecond * 500)
	}
}

func incr(i int) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}
