package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T\n", fun1)
	var p Person
	p.age = 1
	p.name = "aa"
	fmt.Println(p)
	var s1 Shape = p
	fmt.Println(s1.area())

}

func fun1() {
	fmt.Println("are you ok?")
}

type Person struct {
	name string
	age  int
}

func (p Person) area() float64 {
	return 2
}

type Shape interface {
	area() float64
}
