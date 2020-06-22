package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket int = 10
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	// go saleTicket("窗口1")
	// go saleTicket2("窗口2")
	go saleTicket("窗口3")
	go saleTicket2("窗口4")
	wg.Wait()
	time.Sleep(1 * time.Second)
}

func saleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
			fmt.Println(name+"售出一张，还剩", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售完")
			break
		}
		mutex.Unlock()
	}
}

func saleTicket2(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
			fmt.Println(name+"售出一张，还剩", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售完")
			break
		}
		mutex.Unlock()
	}
}
