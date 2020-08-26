package main

import (
	"fmt"
)

func main() {
	defs()
}

func defs() {
	func() {
		fmt.Println("a")
	}()
	func() {
		fmt.Println("b")
	}()
	func() {
		fmt.Println("c")
	}()
	panic("no ok")
}
