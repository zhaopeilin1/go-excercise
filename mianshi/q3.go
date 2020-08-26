package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	func() {
	// 		fmt.Println("A:", i)
	// 		wg.Done()
	// 	}()
	// }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go func(a int) {
			fmt.Println("B:", a)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
