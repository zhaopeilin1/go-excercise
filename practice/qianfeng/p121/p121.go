package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {

	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
	// wg.Add(2)
	// go readDate(1)
	// go readDate(2)
	wg.Add(3)
	go writeData(1)
	go readDate(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main overR")

}

func readDate(i int) {
	defer wg.Done()
	fmt.Println(i, "准备开始读")
	rwMutex.RLock()
	fmt.Println(i, "读锁，正在读")
	time.Sleep(1 * time.Second)
	rwMutex.RUnlock()
	fmt.Println(i, "结束释放读锁")
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "准备开始写:")
	rwMutex.Lock()
	fmt.Println(i, "写锁，正在写")
	time.Sleep(1 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写锁结束")
}
