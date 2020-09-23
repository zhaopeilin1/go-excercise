package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	arr := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}
	fmt.Println(arr)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	ch1 := make(chan string)
	for i := 1; i < 3; i++ {
		go consume("worker"+fmt.Sprint(i), ch1, wg)
	}

	// for _, str := range arr {
	// 	ch1 <- str
	// }
	wg1 := &sync.WaitGroup{}
	wg1.Add(9)
	for i := 1; i < 10; i++ {
		go produce("producer"+fmt.Sprint(i), ch1, wg1)
	}
	wg1.Wait()

	close(ch1)

	wg.Wait()

}

func produce(name string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(500 * time.Millisecond)
	fmt.Println(name)
	ch <- fmt.Sprint(rand.Intn(10))
}

func consume(name string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for str := range ch {
		time.Sleep(1 * time.Second)
		fmt.Println(name, str)
	}
	fmt.Println("close")
}
