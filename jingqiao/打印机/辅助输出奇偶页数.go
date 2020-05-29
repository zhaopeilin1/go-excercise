package main

import "fmt"

func main() {
	a := 49
	for i := 0; i < 10; i++ {
		if i%10 == 9 {
			fmt.Print(a + i*2)
		} else {
			fmt.Print(a+i*2, ",")
		}
	}
	fmt.Println()

}
