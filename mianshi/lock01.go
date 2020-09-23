package main

import (
	"math/rand"
	"sync"

	"time"

	"fmt"
)

var ticket = 10

//var wg = sync.WaitGroup
var lock sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(4)
	go saleTickets("售票员1")
	go saleTickets("售票员2")
	go saleTickets("售票员3")
	go saleTickets("售票员4")
	wg.Wait()

	// time.Sleep(5 * time.Second)

}

func saleTickets(name string) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		lock.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出:", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，出票失败")
			lock.Unlock()

			break
		}
		lock.Unlock()
	}
}
