package main

import (
	"fmt"
)

func main() {
	go printName()
	for i := 0; i < 1000; i++ {
		fmt.Println("main", i)
	}
	fmt.Println("main over")

}

func printName() {
	for i := 0; i < 1000; i++ {
		fmt.Println("goroutin", i)
	}

}
