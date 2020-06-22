package main

import (
	"math/rand"

	"fmt"
	"time"
)

var ticket int = 10

func main() {
	go saleTicket("窗口1")
	go saleTicket("窗口2")
	go saleTicket("窗口3")
	go saleTicket("窗口4")
	time.Sleep(1 * time.Second)
}

func saleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
			fmt.Println(name+"售出一张，还剩", ticket)
			ticket--
		} else {
			fmt.Println(name, "售完")
			break
		}
	}

}
