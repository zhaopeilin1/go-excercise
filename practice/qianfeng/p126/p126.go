package main

import (
	"fmt"
	"time"
)

func main() {
	// t := time.NewTimer(3 * time.Second)
	// fmt.Printf("%T\n", t)
	// fmt.Println(time.Now())

	// ch2 := t.C
	// fmt.Println(<-ch2) // now +3 second

	timer2 := time.NewTimer(5 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 over start")
	}()
	time.Sleep(3 * time.Second)
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stop")
	}

}
