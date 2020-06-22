package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		反射
	*/
	var x int = 3
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("------")
	v := reflect.ValueOf(x)
	fmt.Println("kind", v.Kind(), "type", v.Type())
}
