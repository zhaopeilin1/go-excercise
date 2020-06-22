package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch)
	fmt.Println(time.Now())
	t2 := <-ch
	fmt.Println(t2)
}
